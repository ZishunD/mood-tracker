# 使用官方 Golang 镜像作为构建镜像
FROM golang:1.24-alpine AS build

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum 文件并下载依赖
COPY go.* ./
RUN go mod download

# 复制整个项目的源代码到容器中
COPY . .

# 编译 Go 项目
RUN go build -o main main.go

# 安装必要的依赖（如 ca-certificates）
RUN apk --no-cache add ca-certificates

# 暴露端口
EXPOSE 8080

# 设置容器启动时执行的命令
CMD ["./main"]
