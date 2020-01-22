##### 编译linux 可以执行文件
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o cronpvdata main.go