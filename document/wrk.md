### wrk

- 进入 wsl

  > powershell
  >
  > wsl

- 安装

  - 启动 Ubuntu

    ![ubuntu_run_01](E:\gothmslee\go-junior\document\wsl\ubuntu_run_01.jpg)

    ![ubuntu_run_02](E:\gothmslee\go-junior\document\wsl\ubuntu_run_02.jpg)

  - 验证 Ubuntu 是否成功启动

    > bash：lsb_release -a
    >
    > 输出：
    >
    > No LSB modules are available.
    > Distributor ID: Ubuntu
    > Description:    Ubuntu 22.04.3 LTS
    > Release:        22.04
    > Codename:       jammy

  - 安装 wrk

    - 更新系统

      > sudo apt update && sudo apt upgrade -y

    - 安装依赖

      > sudo apt install -y build-essential git libssl-dev

    - 环境准备

      > 安装 make 工具：sudo apt-get install make
      >
      > 安装 gcc 编译环境：sudo apt-get install build-essential

    - 编译安装wrk：全局代理，系统代理

      >git clone https://github.com/wg/wrk.git
      >cd wrk
      >make 或者 make -j8

    - Error 1：输入 make 命令后

      > /mnt/c/Users/sc/wrk$ make 
      >
      > echo LuaJIT-2.1 
      >
      > LuaJIT-2.1 
      >
      > make: unzip: No such file or directory 
      >
      > make: *** [Makefile:82: obj/LuaJIT-2.1] Error 127

      解决方案：在 Ubuntu 终端中安装 unzip

      > sudo apt update && sudo apt install unzip -y

    - Error 2：

      > /mnt/c/Users/sc/wrk$ make -j8 
      >
      > echo LuaJIT-2.1 
      >
      > LuaJIT-2.1 
      >
      > Archive:  deps/LuaJIT-2.1.zip ec6edc5c39c25e4eb3fca51b753f9995e97215da 
      >
      > LUAJIT src/wrk.lua 
      >
      > /bin/sh: 1: Syntax error: "(" unexpected 
      >
      > make: *** [Makefile:64: obj/bytecode.c] Error 2
      >
      > make: *** Waiting for unfinished jobs....

      解决方案：

      > 在编译前，将 PATH 简化为仅包含 Linux 系统路径，避免 Windows 路径干扰
      >
      > export PATH="/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
      >
      > 重新编译：make clean && make -j8

      原理说明：

      > **PATH 污染问题**：WSL 中默认会将 Windows 的 PATH 合并到 Linux 环境，而 Windows 路径中的空格和括号（如 `Program Files (x86)`）会导致 Shell 解析错误
      >
      > **Makefile 路径处理**：直接引用含特殊字符的路径时，需用引号包裹或转义，否则 Shell 会错误分割命令参数

    - 验证安装

      > ./wrk --version
      >
      > 输出类似 wrk 4.2.0 即为成功

  - 加入环境变量

    > deepseek：wrk 加入环境变量
    >
    > export PATH="$PATH:/mnt/c/Users/sc/wrk"
    >
    > 验证：wrk --version

  - 

- 

