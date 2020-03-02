package template

// Dockerfile 模版
const Dockerfile = `
# 编译

FROM golang:1.14 as build

WORKDIR /build

COPY .  .

# 国内使用的goproxy
#RUN export GOPROXY=https://goproxy.cn

RUN make build_in_docker

# 运行

FROM alpine:latest

WORKDIR /root/

# 调整时区为北京时间
#RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
#    apk add --no-cache ca-certificates tzdata  && \
#    ln -s /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

COPY --from=build /build/{{.project}} .

#EXPOSE <port>

#CMD ["./{{.project}}"]

ENTRYPOINT ["./{{.project}}"]
`
