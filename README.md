##### 程序简介cronpvdata

- 通过gorm访问数据库进行数据的更新
- 设置定时任务进行数据更新
- 利用管道chan功能进行死等
- 同时记录更新记录日志
- 生成随机数


##### 编译linux 可以执行文件
~~~
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o cronpvdata main.go
~~~