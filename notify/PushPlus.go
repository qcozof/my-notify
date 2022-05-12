package notify

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/microcosm-cc/bluemonday"
	"github.com/qcozof/my-notify/global"
)

const maxTitleLen = 100

func PushPlus(title, content string) {

	ruTitle := []rune(title)
	if len(ruTitle) > maxTitleLen {
		title = string(ruTitle[0:maxTitleLen])
	}

	//过滤html
	bm := bluemonday.StripTagsPolicy()
	title = bm.Sanitize(title)

	title = strings.Replace(title, "\n", "", -1)
	title = strings.Replace(title, "\r", "", -1)
	content = strings.Replace(content, "\n", "<br/>", -1)

	config := global.SERVER_CONFIG.PushPlusConfig
	if !config.Enable {
		fmt.Println("PushPlus not enable ! msg will not be sent through PushPlus.")
		return
	}

	remoteUrl := config.ApiUrl
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true} //忽略错误https证书

	//传键值对
	var data = make(map[string]string)
	data["token"] = global.SERVER_CONFIG.PushPlusConfig.Token
	data["title"] = title
	data["content"] = content
	data["template"] = "html"

	btJson, err := json.Marshal(data)
	if err != nil {
		fmt.Println("json.Marshal err :", err)
		return
	}

	req, err := http.NewRequest(http.MethodPost, remoteUrl, bytes.NewBuffer(btJson))

	//设置header
	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	req.Header.Add("Content-Length", strconv.Itoa(len(data)))

	client := &http.Client{}
	resp, err := client.Do(req)

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll err:", err)
	} else {
		fmt.Println(string(b))
		fmt.Println("消息已发送！")
	}

}
