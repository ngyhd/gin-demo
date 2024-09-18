FROM loads/alpine:3.8
#  CGO_ENABLED=0 GOOS=linux  GOARCH=amd64  go  build  main.go
#  docker build -t gin-demo:v1 .
## CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go
###############################################################################
#                                INSTALLATION
###############################################################################

# 设置固定的项目路径
ENV WORKDIR /app/main

# 添加应用可执行文件，并设置执行权限
ADD ./main   $WORKDIR/main
RUN chmod +x $WORKDIR/main


###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
CMD ./main