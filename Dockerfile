# 基于镜像 golang1.5.1
FROM golang:1.5.1

# 代码所运行的目录 
WORKDIR /go/src

# 拷贝文件
ADD ./vendor/src /go/src
ADD ./src/es-cleaner /go/src/es-cleaner

# 设置工作目录
WORKDIR /go/src/es-cleaner

# 编译
RUN go install es-cleaner
