##### 程序简介cronpvdata

- 通过gorm访问数据库进行数据的更新
- 设置定时任务进行数据更新
- 利用管道chan功能进行死等
- 同时记录更新记录日志
- 生成随机数int64

##### 程序说明

1. 该程序只是临时作为程序更新优易电项目中的统计数据
2. 最好的方案，同时创建一套硬件设备齐全的电站
3. 修改统计数据，只是为演示做准备


##### 编译linux 可执行文件

~~~
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o cronpvdata main.go
~~~

end