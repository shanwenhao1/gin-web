# 源镜像
FROM ubuntu:16.04
# 作者
MAINTAINER Razil "swh-email@qq.com"
# 设置工作目录, 因为是编译后的文件, 所以没放在GOPATH下面
WORKDIR $home/test
# 将工程目加入到docker容器中
COPY . $home/test
# 更换软件源, sources.list必须在dockerfile同级目录下
ADD sources.list /etc/apt/
# ENV定义环境变量
# 安装依赖
RUN apt-get install apt-transport-https \
    && apt-get update \
    && apt-get install apt-transport-https \
    && apt-get install libcurl4-openssl-dev libapr1-dev libaprutil1-dev libmxml-dev \
    && wget http://docs-aliyun.cn-hangzhou.oss.aliyun-inc.com/assets/attach/32131/cn_zh/1501595738954/aliyun-oss-c-sdk-3.5.0.tar.gz \
    && wget http://docs-aliyun.cn-hangzhou.oss.aliyun-inc.com/assets/attach/32131/cn_zh/1501595738954/aliyun-oss-c-sdk-3.5.0.tar.gz
# 运行程序, test-app为编译后的go代码
CMD  ["./test-app"]

