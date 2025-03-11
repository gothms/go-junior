### Goland

- Edit Configurations

  - 最常用：package

    ![run_package](E:\gothmslee\go-junior\document\goland\run_package.jpg)

  - 不常用：directory、file

    ![run_directory](E:\gothmslee\go-junior\document\goland\run_directory.jpg)

- 依赖同步：sync dependency
  ![dependency_01](E:\gothmslee\go-junior\document\goland\dependency_01.jpg)

  > Alt+Enter

  ![dependency_02](E:\gothmslee\go-junior\document\goland\dependency_02.jpg)
  
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

- goland 配置 mysql（测试 docker 启动 mysql 成功）
  ![mysql_driver](E:\gothmslee\go-junior\document\goland\mysql_driver.jpg)

  ![mysql_driver_password_root](E:\gothmslee\go-junior\document\goland\mysql_driver_password_root.jpg)

  ![mysql_driver_done](E:\gothmslee\go-junior\document\goland\mysql_driver_done.jpg)

- Linux Edit Configurations

  - linux 命令：GOOS=linux GOARCH=arm go build -o webook .

  - Windows 下 goland 设置：注意 GOARCH=arm

    > 否则编译的文件并不是在 linux 下可执行的

    ![linux_build_setting](E:\gothmslee\go-junior\document\goland\linux_build_setting.png)

  - 设置全局输出路径
    
    > 默认路径：C:\Users\sc\AppData\Local\JetBrains\GoLand2024.3\tmp\GoLand
    
    ![output_directory](E:\gothmslee\go-junior\document\goland\output_directory.jpg)
    
  - 

- 

### npm

- 询问 deepseek 安装及环境配置

- 在 D:\Program Files\nodejs 目录下，创建node_global和node_cache文件夹

- **赋予当前用户对 `D:\Program Files\nodejs` 的完全控制权**：
  
  右键点击 `nodejs` 文件夹 → **属性** → **安全** → 选择当前用户 → 勾选 **“完全控制”** → 应用
  
- **避免将 Node.js 安装在系统目录**：
  下次安装 Node.js 时，选择用户目录（如 `D:\nodejs` 而非 `Program Files`）
  
- 重新安装依赖：Module not found: Can't resolve '@ant-design/cssinjs'
  npm cache clean --force
  npm install：全部
  npm install antd @ant-design/cssinjs --save：针对某一个依赖
  淘宝镜像：npm config set registry https://registry.npmmirror.com/
  
- 

