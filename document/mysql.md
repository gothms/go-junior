### 数据库

- 关联关系
  - 数据库是有状态的服务，横向扩展非常不好，应尽可能减轻数据库的压力
  - 按照DDD的理论，关联关系应该在 Repository 层面处理，而不是 dao 层
  
- 开启 sql 语句的日志

  > db = db.Debug()

- 软删除

  > 业界通用做法
  >
  > 同时，软删除并没有什么用

- 

### MySQL

- 用户名&密码

  > root，1234
  >
  > test，1234

- 使用 Docker Compose 启动数据库
  ![gland_docker_mysql](E:\gothmslee\go-junior\document\picture\gland_docker_mysql.jpg)

  > 原因1：mysql 没启动，docker services 没启动
  >
  > ![goland_docker_mysql_sol](E:\gothmslee\go-junior\document\picture\goland_docker_mysql_sol.jpg)
  >
  > 然后在 services 中启动 docker
  >
  > 原因2：没有科学上网，或配置docker镜像
  >
  > ![gland_docker_mysql_sol_net](E:\gothmslee\go-junior\document\picture\gland_docker_mysql_sol_net.jpg)

- 

### GORM

- 声明模型：https://gorm.io/zh_CN/docs/models.html
- 字段必须大写，否则不会Bind成功，也不会在数据库中创建该字段

### Docker

- 使用 Docker Compose 启动数据库

- Docker Compose 语法

  > https://hub.docker.com/_/mysql
  >
  > https://docs.docker.com/reference/compose-file/

- 安装报错：Docker Desktop requires Windows 10 Pro/Enterprise/Home version 19044 or above.

  > 更新windows异常：你的设备中缺少重要的安全和质量修复
  >
  > 到 microsoft 官网，点击现在更新：https://www.microsoft.com/zh-cn/software-download/windows10

- 运行报错：

  ![docker_desktop](E:\gothmslee\go-junior\document\picture\docker_desktop.png)

  > 原因：没有Linux环境
  >
  > DeepSeek帮忙解决，安装wsl2

- 运行再报错：

  ![wsl2](E:\gothmslee\go-junior\document\picture\wsl2.jpg)

  > 错误信息：
  > deploying WSL2 distributions
  > ensuring main distro is deployed: deploying "docker-desktop": importing WSL distro "当前计算机配置不支持 WSL2。\r\n请启用“虚拟机平台”可选组件，并确保在 BIOS 中启用虚拟化。\r\n通过运行以下命令启用“虚拟机平台”: wsl.exe --install --no-distribution\r\n有关信息，请访问 https://aka.ms/enablevirtualization\r\n错误代码: Wsl/Service/RegisterDistro/CreateVm/HCS/HCS_E_HYPERV_NOT_INSTALLED\r\n" output="docker-desktop": exit code: 4294967295: running WSL command wsl.exe C:\WINDOWS\System32\wsl.exe --import docker-desktop <HOME>\AppData\Local\Docker\wsl\main C:\Program Files\Docker\Docker\resources\wsl\wsl-bootstrap.tar --version 2: 当前计算机配置不支持 WSL2。
  > 请启用“虚拟机平台”可选组件，并确保在 BIOS 中启用虚拟化。
  > 通过运行以下命令启用“虚拟机平台”: wsl.exe --install --no-distribution
  > 有关信息，请访问 https://aka.ms/enablevirtualization
  > 错误代码: Wsl/Service/RegisterDistro/CreateVm/HCS/HCS_E_HYPERV_NOT_INSTALLED
  > : exit status 0xffffffff
  > checking if isocache exists: CreateFile \\wsl$\docker-desktop-data\isocache\: The network name cannot be found.
  >
  > 原因：没有正确安装wsl2，即适用于Linux的Windows子系统
  >
  > WslRegisterDistribution failed with error: 0x80370102
  >
  > Please enable the Virtual Machine Platform Windows feature and ensure virtualization is enabled in the BIOS.

- 安装WSL2

  > 参考：
  > DeepSeek、https://zhuanlan.zhihu.com/p/147233604、https://learn.microsoft.com/zh-cn/windows/wsl/install
  >
  > 需要：卸载360等软件，清理Ubuntu_2204残余，清理VMware或VirtualBox残余，清理注册表

- 验证WSL2

  > wsl -l -v
  >
  > ![wsllv](E:\gothmslee\go-junior\document\picture\wsllv.jpg)
  >
  > wsl：用户名eenee，密码1234，启动wsl
  >
  > > Create a default Unix user account: eenee，回车输入密码
  >
  > uname -a
  >
  > cat /etc/os-release
  >
  > ![wsl_run](E:\gothmslee\go-junior\document\picture\wsl_run.jpg)
  >
  > wsl --shutdown：另起powershell，关闭wsl

- docker compose up

  - [ERROR] [MY-000067] [Server] unknown variable 'default-authentication-plugin=mysql_nativ e_password'.

    > docker-compose.yaml 文件设置为 image: mysql:8.4.4 时，配置不同于 image: mysql:8.0

  -  Cannot downgrade from 80404  to 80041. Downgrade is only permitted between patch releases.

    > 之前是 image: mysql:8.4.4，现在改为 image: mysql:8.0 后，需调整镜像版本
    >
    > 解决方案参考 DeepSeek：docker compose up 命令后报错：Cannot downgrade from 80404  to 80041. Downgrade is only permitted between patch releases.

- docker compose

  > - **Docker Compose 官方文档**
  >   https://docs.docker.com/compose/
  >   包含完整的语法说明、示例和最佳实践。
  >
  > - **Compose Specification**
  >   https://docs.docker.com/compose/compose-file/
  >   详细解释 `docker-compose.yml` 文件的每一个配置项。
  >
  > - **B站视频教程**
  >
  >   **【Docker Compose 从入门到实践】**
  >   [https://www.bilibili.com/video/BV1ZT4y1U7rD](https://search.bilibili.com/all?keyword=docker%20compose)
  >   中文分步骤教程，适合边看边操作。
  >
  > - **博客文章**
  >
  >   **阮一峰的网络日志**
  >   [Docker Compose 教程](http://www.ruanyifeng.com/blog/2018/02/docker-wordpress-tutorial.html)
  >   通过 WordPress 实例讲解 Compose 的用法。

- 

