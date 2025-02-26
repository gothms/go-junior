### Goland

- Edit Configurations

  - 最常用：package

    ![run_package](E:\gothmslee\go-junior\document\operation\run_package.jpg)

  - 不常用：directory、file

    ![run_directory](E:\gothmslee\go-junior\document\operation\run_directory.jpg)

- 依赖同步：sync dependency
  ![dependency_01](E:\gothmslee\go-junior\document\operation\dependency_01.jpg)

  > Alt+Enter

  ![dependency_02](E:\gothmslee\go-junior\document\operation\dependency_02.jpg)
  
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

