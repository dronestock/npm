FROM storezhang/alpine


LABEL author="storezhang<华寅>"
LABEL email="storezhang@gmail.com"
LABEL qq="160290688"
LABEL wechat="storezhang"
LABEL Description="Npm插件，支持发布包到仓库等功能"


# 复制文件
COPY npm /bin



RUN set -ex \
    \
    \
    \
    # 安装NPM相关软件 \
    && apk update \
    && apk add npm \
    && npm install -g npm-cli-adduser \
    \
    \
    \
    # 增加执行权限
    && chmod +x /bin/npm \
    \
    \
    \
    && rm -rf /var/cache/apk/*



ENTRYPOINT /bin/npm
