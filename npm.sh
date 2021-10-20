#!/bin/sh

# 匹配不是Drone插件时的使用配置
[ -z "${PLUGIN_FOLDER}" ] && PLUGIN_FOLDER=${FOLDER}

[ -z "${PLUGIN_REGISTRY}" ] && PLUGIN_REGISTRY=${REGISTRY}
[ -z "${PLUGIN_USERNAME}" ] && PLUGIN_USERNAME=${USERNAME}
[ -z "${PLUGIN_PASSWORD}" ] && PLUGIN_PASSWORD=${PPASSWORD}
[ -z "${PLUGIN_TOKEN}" ] && PLUGIN_TOKEN=${TOKEN}
[ -z "${PLUGIN_EMAIL}" ] && PLUGIN_EMAIL=${EMAIL}

[ -z "${PLUGIN_TAG}" ] && PLUGIN_TAG=${TAG}
[ -z "${PLUGIN_ACCESS}" ] && PLUGIN_ACCESS=${ACCESS}

# 处理默认值
[ -z "${PLUGIN_ACCESS}" ] && PLUGIN_ACCESS="public"


# 处理配置带环境变量的情况
PLUGIN_FOLDER=$(eval echo "${PLUGIN_FOLDER}")

PLUGIN_USERNAME=$(eval echo "${PLUGIN_USERNAME}")
PLUGIN_PASSWORD=$(eval echo "${PLUGIN_PASSWORD}")
PLUGIN_TOKEN=$(eval echo "${PLUGIN_TOKEN}")
PLUGIN_EMAIL=$(eval echo "${PLUGIN_EMAIL}")

PLUGIN_TAG=$(eval echo "${PLUGIN_TAG}")
PLUGIN_ACCESS=$(eval echo "${PLUGIN_ACCESS}")


# 进入目录
cd "${PLUGIN_FOLDER}" || exit

# 取出package.json配置
[ -z "${PLUGIN_TAG}" ] && PLUGIN_TAG=$(yq eval .version package.json --output-format json --prettyPrint)
NAME=$(yq eval .name package.json --output-format json --prettyPrint)

# 初始化NPM授权
npm login --registry "${PLUGIN_REGISTRY}" << EOF
${PLUGIN_USERNAME}
${PLUGIN_PASSWORD}
${PLUGIN_EMAIL}
EOF

# 发布包
EXIST=$(npm view "${NAME}@${PLUGIN_TAG}")
if [ -n "${EXIST}" ]; then
  # 线上存在当前包，撤消发布当前包后再发布
  npm unpublish
  npm version patch
  npm publish --tag "${PLUGIN_TAG}" --access "${PLUGIN_ACCESS}"
else
  # 直接发布
  npm publish --tag "${PLUGIN_TAG}" --access "${PLUGIN_ACCESS}"
fi
