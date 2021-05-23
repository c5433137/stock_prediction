FROM debian:buster-slim

RUN mkdir -p /app
RUN mkdir -p /app/log
WORKDIR /app
# 将构建好的二进制文件拷贝进镜像
COPY  ./stock_prediction /app/stock_prediction
EXPOSE 8080
# 启动 Web 服务
CMD ["/app/stock_prediction"]