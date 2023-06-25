package notify

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/qcozof/my-notify/global"
	"github.com/spf13/viper"
)

func InitConfig(config string) {
	if len(config) == 0 {
		//dir, _ := os.Getwd()
		config = "../config.yaml"
	}

	v := viper.New()
	v.SetConfigFile(config)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.SERVER_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.SERVER_CONFIG); err != nil {
		fmt.Println(err)
	}
	global.VIPER = v
}

func NotifyAll(title, message string, token ...string) {
	go PushPlus(title, message, token...)
	go Telegram(message)
	go Discord(title, message)

	var email = Email{
		Subject: title,
		Body:    message,
	}

	var dingDing = DingDing{
		Content: message,
	}

	var chinaMobileSms = ChinaMobileSms{
		TemplateId: title,
		Params:     message,
	}

	var chinaUnicomSms = ChinaUnicomSms{
		TemplateCode:      title,
		TemplateParamJson: message,
	}

	go email.Send()
	go dingDing.Send()
	go chinaMobileSms.Send()
	go chinaUnicomSms.Send()
}
