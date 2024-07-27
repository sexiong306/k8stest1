# 基础镜像
FROM alpine:latest

# 更换 apk 源
RUN sed -i 's!http://dl-cdn.alpinelinux.org/!https://mirrors.tencent.com/!g' /etc/apk/repositories

# 安装 apk
RUN set -eux && \
    apk add bind-tools && \
    apk add busybox-extras && \
    apk add findutils && \
    apk add tcpdump && \
    apk add tzdata && \
    apk add curl && \
    apk add curl && \apk add bash && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone && \
    date

# 制作者
MAINTAINER aibum

# 切换工作目录
WORKDIR /usr/local/app

RUN mkdir /usr/local/app -p
RUN mkdir /usr/local/app -p
RUN mkdir /usr/local/app/bin -p
RUN mkdir /usr/local/app/conf -p
RUN mkdir /usr/local/app/log -p
RUN mkdir /usr/local/app/data -p

# 将当前目录下的所有文件添加到WorkDIR目录中
COPY ./scripts/start.sh /usr/local/app/
COPY ./helloworld /usr/local/app/bin/
COPY ./trpc_go.yaml /usr/local/app/bin/

RUN chmod 777 /usr/local/app/start.sh
RUN chmod 777 /usr/local/app/bin/helloworld
RUN chmod 777 /usr/local/app/bin/trpc_go.yaml

# 运行启动脚本
CMD ["/bin/bash", "start.sh"]
