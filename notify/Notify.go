package notify

import (
	_ "github.com/qcozof/my-notify/initialize"
)

func NotifyAll(title ,message string){
	go PushPlus(title, message)
	go Telegram(message)
	go Discord(title,message)
}