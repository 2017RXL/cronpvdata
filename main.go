package main

import (
	"cronpvdata/model"
	"github.com/robfig/cron"
	"log"
)

func main() {
	model.Logger.Info("定时修改数据任务已经启动！！")
	go cronDoWork()
	var str chan string
	d := <-str
	log.Println(d)

}
func cronDoWork() {
	c := cron.New()
	//每个1分钟执行一次
	//c.AddFunc("0 */1 * * * ?", func() {
	c.AddFunc("0 0 2 * * ?", func() {
		model.Logger.Info("开始更新数据了")
		model.UpdateData()
		model.Logger.Info("更新数据完毕了")
	})
	c.Start()
}
