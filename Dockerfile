FROM storezhang/alpine AS builder


# Yaml修改程序版本
ENV YQ_VERSION 4.13.4
ENV YQ_BINARY yq_linux_amd64


RUN wget https://ghproxy.com/https://github.com/mikefarah/yq/releases/download/v${YQ_VERSION}/${YQ_BINARY} --output-document /usr/bin/yq
RUN chmod +x /usr/bin/yq



# 打包真正的镜像
FROM storezhang/alpine


MAINTAINER storezhang "storezhang@gmail.com"
LABEL architecture="AMD64/x86_64" version="latest" build="2021-10-20"
LABEL Description="NPM发布包插件"


# 复制文件
COPY --from=builder /usr/bin/yq /usr/bin/yq
COPY npm.sh /bin



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
    && chmod +x /bin/npm.sh \
    \
    \
    \
    && rm -rf /var/cache/apk/*



ENTRYPOINT /bin/npm.sh
