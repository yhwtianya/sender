package http

import (
	"fmt"
	"net/http"

	"github.com/open-falcon/sender/proc"
)

// 获得sms,mail发送总数量
func configProcRoutes() {

	http.HandleFunc("/count", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("sms:%v, mail:%v", proc.GetSmsCount(), proc.GetMailCount())))
	})

}
