package proc

import (
	"sync/atomic"
)

var smsCount, mailCount uint32

// 获取sms发送数量
func GetSmsCount() uint32 {
	return atomic.LoadUint32(&smsCount)
}

// 获取mail发送数量
func GetMailCount() uint32 {
	return atomic.LoadUint32(&mailCount)
}

// 增加sms发送数量
func IncreSmsCount() {
	atomic.AddUint32(&smsCount, 1)
}

// 增加mail发送数量
func IncreMailCount() {
	atomic.AddUint32(&mailCount, 1)
}
