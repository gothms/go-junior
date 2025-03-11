### Docker

- Enable Kubernetes

  - 完全卸载Docker，清理注册表
  - 卸载防护软件，如360，关闭windows防火墙
  - 端口可能被占用
  - 开启vpn代理，并最好使用全局代理模式

  ![enable_kubernetes](E:\gothmslee\go-junior\document\docker\enable_kubernetes.jpg)

  ![enable_kubernetes_done](E:\gothmslee\go-junior\document\docker\enable_kubernetes_done.jpg)

- 

### K8s部署web服务器

#### 编译

- 编译在 linux 下运行的可执行文件

  - mac 命令：GOOS=linux GOARCH=arm go build -o webook .

  - window 命令：

    ```
    set GOOS=linux
    set GOARCH=arm
    go build -o webook .
    
  - 以上命令编译的文件不是 GOOS=linux GOARCH=arm，导致 pods 启动不成功

    - Error 1：

      > CrashLoopBackOff：Back-off restarting failed container webook-record in pod ......

    - Error 2：

      > docker run -it --rm gothms/webook:v0.0.1 
      >
      > exec /app/webook: no such file or directory

    - 解决方案：

      > 1.command: ["/bin/bash", "-ce", "tail -f /dev/null"]
      >
      > 
      > 这会导致容器只执行挂起命令，实际应用并未启动，因此 Service 和 Ingress 无法将流量转发到有效的后端，直接引发 502 错误（配置 Ingress 后）
      > 
      >
      > 2.参考 goland.md

- 查看 wsl 安装的 ubuntu 的版本

  > 打开 wsl.exe，输入 lsb_release -a

- 打包成一个镜像，且 Error

  > docker build -t gothms/webook:v0.0.1 .
  >
  > 删除镜像：docker rmi -f gothms/webook:v0.0.1
  >
  > 本地运行测试：docker run -it --rm gothms/webook:v0.0.1

  > ERROR: failed to solve: ubuntu:24.04: failed to resolve source metadata for docker.io/library/ubuntu:24.04: failed to authorize: failed to fetch anonymous token: Get "https://auth.docker.io/token?scope=repository%3Alibrary%2Fubuntu%3Apull&service=registry.docker.io": dial tcp 104.244.46.71:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.

  ![docker_build_v0001](E:\gothmslee\go-junior\document\docker\docker_build_v0001.jpg)

  > 解决方式一：关掉 vpn，并配置代理
  >
  > "registry-mirrors": [
  > "https://registry.cn-hangzhou.aliyuncs.com",
  > "https://docker.mirrors.ustc.edu.cn",
  > "https://docker.1ms.run",
  > "https://docker.xuanyuan.me",
  > "https://mirror.ccs.tencentyun.com",
  > "https://registry.docker-cn.com"
  > ]
  >
  > {
  > "registry-mirrors": ["https://docker.mirrors.ustc.edu.cn"]
  > }
  >
  > 方式二：https://xixibeiu6y.xyz/user
  >
  > 开启 Tun 模式

- Makefile：只有 mac 系统可以

  ![docker_build_v0001](E:\gothmslee\go-junior\document\docker\docker_build_v0001.jpg)

#### 运行 Deployment

- kubectl apply -f webook-deployment.yaml

- kubectl get pods

  >kubectl get pods |grep record
  >
  >windows：kubectl get pods |select-string record

- kubectl get deployment |select-string record

#### 运行 Service

- kubectl apply -f webook-service.yaml
- kubectl get pods
- kubectl get deployment

#### kubectl delete

- kubectl delete deployment --all
- kubectl delete service --all
- kubectl delete pods --all
- kubectl delete pvc --all
- kubectl delete pv --all

#### kubectl logs

- kubectl logs webook-record-6569859569-5l7lp

#### Error

- 描述

  > PS E:\gothmslee\go-junior\webook> kubectl get pods
  > NAME                                     READY   STATUS   RESTARTS      AGE
  > webook-record-service-6569859569-4pjf4   0/1     Error    3 (35s ago)   28m
  > webook-record-service-6569859569-4qrvg   0/1     Error    4 (55s ago)   28m
  > webook-record-service-6569859569-jm2p7   0/1     Error    2 (24s ago)   27s
  > PS E:\gothmslee\go-junior\webook> kubectl get deployment
  > NAME                    READY   UP-TO-DATE   AVAILABLE   AGE
  > webook-record-service   0/3     3            0           29m

- 查看 Pod 的详细信息：kubectl describe pod <pod-name>

  > Failed to pull image "gothms/webook:v0.0.1": Error response from
  >  daemon: failed to resolve reference "docker.io/gothms/webook:v0.0.1": unexpected status from HEAD request to https://docker.1ms.run/v2/gothms/webook/manifests/v0.0.1?ns=docker.io: 500 Internal Server Error

- 解决方案

  > https://developer.aliyun.com/article/1354807
  >
  > 
  > command: ["/bin/bash", "-ce", "tail -f /dev/null"]
  > 

- 

### K8s 部署 MySQL

- kubectl get service

  ![docker-k8s-mysql-test](E:\gothmslee\go-junior\document\docker\docker-k8s-mysql-test.jpg)

  ![docker-k8s-mysql-test-done](E:\gothmslee\go-junior\document\docker\docker-k8s-mysql-test-done.jpg)

- Error

  > kubelet            Error: ImagePullBackOff
  >
  > Failed to pull image "mysql:8.0": Error response from daemon: 
  > failed to resolve reference "docker.io/library/mysql:8.0": failed to do request: Head "https://registry-1.docker.io/v2/library/mysql/manifests/8.0": EOF

  - 测试：docker --debug pull mysql:8.0

  - 方案："registry-mirrors": ["https://docker.mirrors.ustc.edu.cn"]，没完全成功

  - 详细方案：

    > 卸载 360，清理注册表
    >
    > 关闭 windows 防火墙
    >
    > XiXicats，全局，系统代理

- 

### K8s 部署 Redis

