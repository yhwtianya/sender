package cron

import (
	"github.com/open-falcon/sender/g"
)

var (
	// 控制调用sms接口并发数
	SmsWorkerChan chan int
	// 控制调用mail接口并发数
	MailWorkerChan chan int
)

// 初始化sms和mail并发数量
func InitWorker() {
	workerConfig := g.Config().Worker
	SmsWorkerChan = make(chan int, workerConfig.Sms)
	MailWorkerChan = make(chan int, workerConfig.Mail)
}
