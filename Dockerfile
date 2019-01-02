# 源镜像
FROM ubuntu:16.04
# 作者
MAINTAINER Razil "swh-email@qq.com"
# 设置工作目录, 因为是编译后的文件, 所以没放在GOPATH下面
WORKDIR /root/oss-server
# 将工程目录加入到docker容器中
COPY . /root/oss-server
# 更换软件源, sources.list必须在dockerfile同级目录下
# ADD sources.list /etc/apt/
# 添加阿里云依赖
RUN echo "#添加阿里源\n\
deb http://mirrors.aliyun.com/ubuntu/ bionic main restricted universe multiverse\n\
deb http://mirrors.aliyun.com/ubuntu/ bionic-security main restricted universe multiverse\n\
deb http://mirrors.aliyun.com/ubuntu/ bionic-updates main restricted universe multiverse\n\
deb http://mirrors.aliyun.com/ubuntu/ bionic-proposed main restricted universe multiverse\n\
deb http://mirrors.aliyun.com/ubuntu/ bionic-backports main restricted universe multiverse\n\
deb-src http://mirrors.aliyun.com/ubuntu/ bionic main restricted universe multiverse\n\
deb-src http://mirrors.aliyun.com/ubuntu/ bionic-security main restricted universe multiverse\n\
deb-src http://mirrors.aliyun.com/ubuntu/ bionic-updates main restricted universe multiverse\n\
deb-src http://mirrors.aliyun.com/ubuntu/ bionic-proposed main restricted universe multiverse\n\
deb-src http://mirrors.aliyun.com/ubuntu/ bionic-backports main restricted universe multiverse">>/etc/apt/sources.list
# ENV定义环境变量
# 指定行添加
    # sed -i "1i\!/bin/bash" $HOME/.bashrc
# 安装依赖
RUN apt-get update \
    && DEBIAN_FRONTEND=noninteractive apt-get -y install libcurl4-openssl-dev libapr1-dev libaprutil1-dev libmxml-dev \
    && DEBIAN_FRONTEND=noninteractive echo n | dpkg-reconfigure dash \
    && echo -e "export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/root/oss-server/kw_media/lib:/usr/local/lib\n\
export C_INCLUDE_PATH=$C_INCLUDE_PATH:/usr/include/apr-1.0">>$HOME/.bashrc && source $HOME/.bashrc \
    && DEBIAN_FRONTEND=noninteractive apt-get -y install wget git tar cmake gcc g++ vim net-tools lrzsz \
    && wget http://docs-aliyun.cn-hangzhou.oss.aliyun-inc.com/assets/attach/32131/cn_zh/1501595738954/aliyun-oss-c-sdk-3.5.0.tar.gz -P $HOME/ \
    && git clone -b fix-sample-bug https://github.com/aliyun/aliyun-media-c-sdk.git $HOME/oss-media-c \
    && find $HOME/oss-media-c/sample/*.c -exec sed -i "s/sleep(/\/\/sleep(/g" {} \; && find $HOME/oss-media-c/sample/*.h -exec sed -i "s/sleep(/\/\/sleep(/g" {} \; \
    && find $HOME/oss-media-c/test/*.c -exec sed -i "s/sleep(/\/\/sleep(/g" {} \; && find $HOME/oss-media-c/test/*.h -exec sed -i "s/sleep(/\/\/sleep(/g" {} \; \
    && sed -i '1i\#include "src/oss_media_client.h"' $HOME/oss-media-c/test/test_all.c \
    && tar -zxvf $HOME/aliyun-oss-c-sdk-3.5.0.tar.gz -C $HOME/ \
    && cd $HOME/aliyun-oss-c-sdk-3.5.0 && cmake . && make && make install \
    && cd $HOME/oss-media-c/src && sed -i "s/int OSS_MEDIA_AAC_SAMPLE_RATE = 1024;/int OSS_MEDIA_AAC_SAMPLE_RATE = 512;/g" $HOME/oss-media-c/src/oss_media_define.c \
    && cd $HOME/oss-media-c && cmake . && make && make install \
    && cd /root/oss-server/ && chmod +777 oss-server
# 运行程序, oss-server为编译后的go代码
CMD  ["./oss-server"]

