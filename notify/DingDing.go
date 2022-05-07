/**
使用示例：
	var dd = notify.DingDing{
		ApiUrl: "https://oapi.dingtalk.com/robot/send?access_token=xxxxx",
		Content: "xx接口异常通知测试",
	}
	_,err := dd.Send()
 */
package notify

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type DingDing struct {
	ApiUrl  string //接口地址
	Content string //要发送的内容
	data    data   //参数结构体
}

//发送消息结构体
type data struct {
	MsgType             string              `json:"msgtype"`
	DingDingTextContent DingDingTextContent `json:"text"`
}

type DingDingTextContent struct {
	Content string `json:"content"`
}

//返回结果
type dingDingResult struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

/*curl --location --request POST 'https://oapi.dingtalk.com/robot/send?access_token='$token \
--header 'Content-Type: application/json' \
--data "{
\"msgtype\": \"text\",
\"text\": {
\"Content\": \"${Content}\"
}
}"*/
func (srv *DingDing) Send() (res string, err error) {

	var remoteUrl = srv.ApiUrl
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true} //忽略错误https证书

	srv.data.MsgType = "text"
	srv.data.DingDingTextContent.Content = srv.Content

	jsonParams, err := json.Marshal(&(srv.data))
	if err != nil {
		return
	}
	fmt.Println(string(jsonParams))

	rdContent := bytes.NewReader(jsonParams)
	req, err := http.NewRequest(http.MethodPost, remoteUrl, rdContent)

	//设置header
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(rdContent.Len()))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var ddResult dingDingResult
	err = json.Unmarshal(body, &ddResult)
	if err != nil {
		return
	}

	//返回正确结果
	if ddResult.ErrCode == 0 {
		res = string(body)
		return
	}

	err = errors.New(ddResult.ErrMsg + ":" + ddResult.ErrMsg)
	return
}
