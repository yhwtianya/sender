package model

import (
	"fmt"
)

// alarm写入sms队列的数据结构
type Sms struct {
	Tos     string `json:"tos"`
	Content string `json:"content"`
}

// alarm写入mail队列的数据结构
type Mail struct {
	Tos     string `json:"tos"`
	Subject string `json:"subject"`
	Content string `json:"content"`
}

func (this *Sms) String() string {
	return fmt.Sprintf(
		"<Tos:%s, Content:%s>",
		this.Tos,
		this.Content,
	)
}

func (this *Mail) String() string {
	return fmt.Sprintf(
		"<Tos:%s, Subject:%s, Content:%s>",
		this.Tos,
		this.Subject,
		this.Content,
	)
}
