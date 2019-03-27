package cron

import (
	"log"
	"time"

	"github.com/open-falcon/sender/g"
	"github.com/open-falcon/sender/model"
	"github.com/open-falcon/sender/proc"
	"github.com/open-falcon/sender/redis"
	"github.com/toolkits/net/httplib"
)

// 持续读取sms队列数据，并发调用发送sms接口
func ConsumeSms() {
	queue := g.Config().Queue.Sms
	for {
		// 无限循环
		L := redis.PopAllSms(queue)
		if len(L) == 0 {
			time.Sleep(time.Millisecond * 200)
			continue
		}
		SendSmsList(L)
	}
}

// 并发调用发送Sms接口
func SendSmsList(L []*model.Sms) {
	for _, sms := range L {
		SmsWorkerChan <- 1
		go SendSms(sms)
	}
}

// 调用发送Sms接口
func SendSms(sms *model.Sms) {
	defer func() {
		<-SmsWorkerChan
	}()

	url := g.Config().Api.Sms
	r := httplib.Post(url).SetTimeout(5*time.Second, 2*time.Minute)
	r.Param("tos", sms.Tos)
	r.Param("content", sms.Content)
	resp, err := r.String()
	if err != nil {
		log.Println(err)
	}

	proc.IncreSmsCount()

	if g.Config().Debug {
		log.Println("==sms==>>>>", sms)
		log.Println("<<<<==sms==", resp)
	}

}
