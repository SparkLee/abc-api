FROM golang:1.23 AS builder

COPY . /src
WORKDIR /src

RUN GOPROXY=https://goproxy.cn make build

FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/bin /app

# 添加默认项目文件
COPY ./configs/config.yaml /data/conf/
#VOLUME /data/conf

WORKDIR /app

EXPOSE 8000
EXPOSE 9000

CMD ["./abc", "-conf", "/data/conf"]
