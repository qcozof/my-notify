/**
 * @Description telegram notify
 **/
package notify

import (
	"encoding/json"
	"fmt"
	"github.com/qcozof/my-notify/global"
	"github.com/qcozof/my-notify/model"
	"io/ioutil"
	"net/http"
	"net/url"
)

func Telegram(text string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Telegram err:",err)
		}
	}()

	config := global.SERVER_CONFIG.TelegramConfig
	if !config.Enable{
		fmt.Println("Telegram not enable !")
		return
	}

	remoteUrl := fmt.Sprintf("%s/bot%s/sendMessage",config.ApiUrl,config.Token)
	//http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true} //忽略错误https证书

	chatId :=config.ChatId
	values := url.Values{
		"chat_id":    []string{chatId},
		"text":       []string{text},
		"parse_mode": []string{"html"},
	}
	resp, err := http.PostForm(remoteUrl, values)

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll err:", err)
	} else {
		var resModel model.TelegramRespModel
		err := json.Unmarshal(b,&resModel)
		if err !=nil{
			fmt.Println(err)
			return
		}

		if resModel.Ok {
			fmt.Println(fmt.Sprintf("telegram消息已发送！chat_id:%s",chatId))
			return
		}
		fmt.Println(fmt.Sprintf("telegram消息发送失败%d:%s;chat_id:%s",resModel.ErrorCode,resModel.Description,chatId))

	}

}