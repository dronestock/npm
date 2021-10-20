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
VERSION=$(yq eval .version package.json --output-format json --prettyPrint | xargs)
NAME=$(yq eval .name package.json --output-format json --prettyPrint | xargs)
PLUGIN_REGISTRY=$(yq eval .publishConfig.registry package.json --output-format json --prettyPrint | xargs)

# 处理默认值
[ -z "${PLUGIN_REGISTRY}" ] && PLUGIN_REGISTRY="https://registry.npmjs.org/"
[ -z "${PLUGIN_ACCESS}" ] && PLUGIN_ACCESS="public"

# 打印相关参数
cat<<EOF
包：${NAME}
版本：${VERSION}
仓库：${PLUGIN_REGISTRY}
EOF

# 初始化NPM授权
if [ -n "${PLUGIN_TOKEN}" ]; then
  echo "使用Token登录"
  URL=$(echo ${PLUGIN_REGISTRY} | sed 's/https\?:\/\///')
  npm config set //"${URL}":_authToken "${PLUGIN_TOKEN}"
else
  echo "使用用户名密码登录"
  npm-cli-adduser --username "${PLUGIN_USERNAME}" --password "${PLUGIN_PASSWORD}" --email "${PLUGIN_EMAIL}" --registry ${PLUGIN_REGISTRY}
fi
npm config set registry ${PLUGIN_REGISTRY}
npm config set always-auth true
npm config set strict-ssl false

# 发布包
PACKAGE_NAME="${NAME}@${VERSION}"
EXIST=$(npm view "${PACKAGE_NAME}")
COMMAND="npm publish"
if [ -n "${PLUGIN_ACCESS}" ]; then
  COMMAND="${COMMAND} --access ${PLUGIN_ACCESS}"
fi
if [ -n "${PLUGIN_TAG}" ]; then
  COMMAND="${COMMAND} --tag ${PLUGIN_TAG}"
fi

if [ -n "${EXIST}" ]; then
  echo "线上存在相同的版本【${VERSION}】，当前版本未发布"
else
  echo "发布包到：${PLUGIN_REGISTRY}${PLUGIN_USERNAME}/${PACKAGE_NAME}"
  ${COMMAND}
fi
