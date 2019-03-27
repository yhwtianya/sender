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

// 持续读取mail队列数据，并发调用发送mail接口
func ConsumeMail() {
	queue := g.Config().Queue.Mail
	for {
		L := redis.PopAllMail(queue)
		if len(L) == 0 {
			time.Sleep(time.Millisecond * 200)
			continue
		}
		SendMailList(L)
	}
}

// 并发调用发送mail接口
func SendMailList(L []*model.Mail) {
	for _, mail := range L {
		MailWorkerChan <- 1
		go SendMail(mail)
	}
}

// 调用发送mail接口
func SendMail(mail *model.Mail) {
	defer func() {
		<-MailWorkerChan
	}()

	url := g.Config().Api.Mail
	r := httplib.Post(url).SetTimeout(5*time.Second, 2*time.Minute)
	r.Param("tos", mail.Tos)
	r.Param("subject", mail.Subject)
	r.Param("content", mail.Content)
	resp, err := r.String()
	if err != nil {
		log.Println(err)
	}

	proc.IncreMailCount()

	if g.Config().Debug {
		log.Println("==mail==>>>>", mail)
		log.Println("<<<<==mail==", resp)
	}

}
