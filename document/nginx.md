### Nginx

- 安装 helm

  > 第一步：安装 helm 到 windows 本地
  >
  > 第二步：使用 helm 安装 ingress-nginx

  - Linux 或类 Unix 系统：

    > curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3
    > chmod 700 get_helm.sh
    > ./get_helm.sh

  - Windows 系统：

    > curl.exe -fsSL -o get_helm.ps1 https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3
    >
    > chmod 700 get_helm.sh：通过 Everything 找到 get_helm.sh 文件，然后 Git Bash Here，再运行此命令
    >
    > ./get_helm.sh：git 窗口运行此命令

  - Linux 或类 Unix 系统：

    > helm upgrade --install ingress-nginx ingress-nginx \
    > --repo https://kubernetes.github.io/ingress-nginx \
    > --namespace ingress-nginx --create-namespace

  - Windows 系统：

    > helm upgrade --install ingress-nginx ingress-nginx `
    > --repo https://kubernetes.github.io/ingress-nginx `
    > --namespace ingress-nginx --create-namespace
    
  - DeepSeek 方案

    - 以 **管理员身份** 打开 PowerShell，执行：

      >Set-ExecutionPolicy Bypass -Scope Process -Force
      >[System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072
      >iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))

    - 安装 Helm

      > choco install kubernetes-helm -y

    - 验证安装

      > helm version
      >
      > 应输出类似：version.BuildInfo{Version:"v3.14.0", ...}
      
    - **安装 Nginx Ingress Controller**

      >helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
      >helm repo update
      >helm upgrade --install ingress-nginx ingress-nginx/ingress-nginx --namespace ingress-nginx --create-namespace

      ![docker-k8s-nginx-install](E:\gothmslee\go-junior\document\docker\docker-k8s-nginx-install.jpg)

- 
