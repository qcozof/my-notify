package notify

import (
	_ "github.com/qcozof/my-notify/initialize"
)

func NotifyAll(title ,message string){
	PushPlus(title, message)
	Telegram(message)
	Discord(title,message)
}