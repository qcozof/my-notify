package notify

import (
	_ "github.com/qcozof/my-notify/initialize"
)

func NotifyAll(title ,message,token string){
	go PushPlus(title, message,token)
	go Telegram(message)
	go Discord(title,message)
}