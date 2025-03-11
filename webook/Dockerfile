FROM ubuntu:24.04
# 安装依赖库（根据实际情况调整）
RUN apt-get update && apt-get install -y libc6
# 复制可执行文件并赋予权限
COPY webook /app/webook
RUN chmod +x /app/webook  # 赋予可执行权限
WORKDIR /app
CMD ["/app/webook"]
