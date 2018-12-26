# Docker

[Docker book](docker.pdf)

## 基本概念

属于操作系统层面的虚拟化技术. 相对于传统虚拟机技术并不虚拟出一套硬件, 容器的应用进程直接运行于宿主的内核, 因此更为轻便.

Docker的三个基本概念:
- 镜像(Image)
    - 镜像相当于Linux的[root文件系统](https://blog.csdn.net/zuidao3105/article/details/79386874)
    - 分层存储: 镜像构建会一层层构建, 前一层是后一层的基础. 当前层的修改并不会实际影响上一层
    (比如删除上一层文件只是在当前层标记, 上层文件实际上还是跟随着镜像)
- 容器(Container)
    - 相当于镜像的实例(以面向对象思想). 容器的实质是进程, 运行于属于自己的命名空间.因此可以拥有自己的root文件系统、
    网络配置等.
    - 分层存储, 以镜像为基础层, 在其上创建当前容器的存储层.容器储存层的生命周期和容器一样.
        - 容器存储层应保持无状态化, 容器不应向其存储层写入任何数据. 而是应使用`数据卷(Volume)`或绑定宿主目录, 
        以跳过容器存储层的读写.
- 仓库(Repository): 集中的存储、分发镜像的服务仓库
    - 仓库内以标签区分镜像, 用<仓库名>:<标签>指定特定镜像(标签不给出则默认为latest). 示例: ubuntu:14.04
    

## 安装与使用

- 安装内核可选模块
    ```bash
    sudo apt-get update
    sudo apt-get install linux-image-extra-$(uname -r) linux-image-extra-virtual
    ```
- 使用国内源下载
    - 添加HTTPS包和CA证书
    ```bash
    sudo apt-get install apt-transport-https ca-certificates curl software-properties-common
    ```
    - 为确保下载软件包的合法性, 需添加软件源的GPG密钥
    ```bash
    curl -fsSL https://mirrors.aliyun.com/docker-ce/linux/ubuntu/gpg | sudo apt-key add -
    ```
    - 在source.list中添加Docker软件源
    ```bash
    sudo add-apt-repository "deb [arch=amd64] https://mirrors.aliyun.com/docker-ce/linux/ubuntu $(lsb_release -cs) stable"
    sudo apt-get update
    ```
- 安装Docker CE
    ```bash
    sudo apt-get install docker-ce
    ```
   - 使用脚本自动安装(注意如果安装过Docker, 脚本执行会出错)
   ```bash
   curl -fsSL get.docker.com -o get-docker.sh
   sudo sh get-docker.sh --mirror Aliyun
   ```
- 启动Docker CE
    ```bash
    sudo systemctl enable docker
    sudo systemctl start docker
    # ubuntu 14.04则使用以下命令启动  
    sudo service docker start
    ```
- 创建用户组: 出于安全考虑, 一般使用`docker`组的用户通过Unix socket与Docker引擎通讯, 而不是使用root用户.
    ```bash
    sudo groupadd docker
    ```