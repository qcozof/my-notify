/**
* @Description:一信通平台发短信
 */
package notify

/*
   例（无账号，未测试）：
	var dd = notify.ChinaUnicomSms{
		PhoneNumberJson: "[\"1333333333\"]",
		TemplateParamJson: "",
	}
	res,err := dd.Send()
	fmt.Println(res,err)

*/

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/qcozof/my-notify/config"
	"github.com/qcozof/my-notify/global"
	"github.com/qcozof/my-notify/utils"

	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

type ChinaUnicomSms struct {
	PhoneNumberJson   string //电话号码 ["1590000****","1350000****"]
	TemplateCode      string //短信模板CODE。如 SMS_152550005
	TemplateParamJson string //短信模板变量对应的实际值，JSON格式。 [{"name":"TemplateParamJson"},{"name":"TemplateParamJson"}]
}

type KV map[string]string

//发送短信, templateCode模板编码(默认调用配置中的)
//如果短信模板没有备案，则http状态码为505
//接口有限制调用方IP
func (srv *ChinaUnicomSms) Send(templateCode string) (aliSmsResult *dysmsapi20170525.SendBatchSmsResponse, err error) {
	conf := global.SERVER_CONFIG.UmsConfig
	if len(templateCode) == 0 {
		templateCode = conf.TemplateCode
	}
	params, serialNumber, err := srv.buildParams(templateCode, conf)
	if err != nil {
		return
	}

	resp, err := srv.send(conf.Url, params, time.Duration(10)*time.Second)
	if err != nil {
		fmt.Printf("一信通短信发送失败,单号:%d,错误信息:%s\n", serialNumber, err.Error())
		return
	}

	aliSmsResult, err = srv.dealResponse(resp, serialNumber)
	return
}

func (srv *ChinaUnicomSms) buildParams(templateCode string, conf config.UmsConfig) (io.Reader, int, error) {
	serialNumber := int(time.Now().UnixNano())

	phones := strings.Replace(srv.PhoneNumberJson, "\"", "", -1)
	phones = strings.Replace(phones, "[", "", -1)
	phones = strings.Replace(phones, "]", "", -1)
/*	template, err := srv.mappingTemplate(templateCode)
	if err != nil {
		return nil, 0, err
	}
	content, err := srv.buildContent(template)*/
	content, err := srv.buildContent(templateCode)
	if err != nil {
		return nil, 0, err
	}
	params := "?SpCode=" + conf.SPCode + "&LoginName=" + conf.LoginName + "&Password=" + conf.Password + "&MessageContent=" + content + "&UserNumber=" + phones + "&SerialNumber=" + strconv.Itoa(serialNumber) + "&ScheduleTime=&f=1"

	fmt.Println("发送一信通短信,模板:" + templateCode + ",参数:" + params)
	return transform.NewReader(bytes.NewReader([]byte(params)), simplifiedchinese.GBK.NewEncoder()), serialNumber, nil
}

func (srv *ChinaUnicomSms) unmarshalParamsJson() (params KV, err error) {
	if len(srv.TemplateParamJson) == 0 {
		err = errors.New("缺失templateParamJson参数")
		return
	}
	var param []KV
	err = json.Unmarshal([]byte(srv.TemplateParamJson), &param)
	if err != nil {
		return
	}

	var p KV = make(KV)
	for _, kv := range param {
		for key := range kv {
			p[key] = kv[key]
		}
	}

	return p, nil
}

func (srv *ChinaUnicomSms) buildContent(template string) (string, error) {
	if ok, _ := regexp.MatchString(`\$\{[\w]+\}`, template); ok {
		params, err := srv.unmarshalParamsJson()
		if err != nil {
			return "", err
		}
		re, _ := regexp.Compile(`\$\{[\w]+\}`)
		sub := re.FindAllStringSubmatch(template, -1)

		for _, v := range sub {
			key := strings.Replace(v[0], "${", "", 1)
			key = strings.Replace(key, "}", "", 1)

			if _, ok := params[key]; !ok {
				return "", errors.New("templateParamJson缺少必要的key:" + key)
			}

			template = strings.Replace(template, v[0], params[key], 1)
		}
	}

	return template, nil
}

func (srv *ChinaUnicomSms) send(url string, params io.Reader, timeout time.Duration) (*http.Response, error) {
	paramsStr, _ := ioutil.ReadAll(params)
	req, err := http.NewRequest("GET", url+string(paramsStr), nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	client.Timeout = timeout
	resp, err := client.Do(req)
	return resp, err
}

func (srv *ChinaUnicomSms) dealResponse(resp *http.Response, serialNumber int) (aliSmsResult *dysmsapi20170525.SendBatchSmsResponse, err error) {
	body, err := ioutil.ReadAll(transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder()))
	if err != nil {
		fmt.Printf("一信通短信发送失败,单号:%d,错误信息:%s\n", serialNumber, err.Error())
		return
	}

	fmt.Printf("一信通短信发送完成,单号:%d,返回值:%s,状态码:%d\n", serialNumber, body, resp.StatusCode)


	queryUrl := "?"+string(body)
	code,err := utils.GetValFromQuery(queryUrl,"result")
	if code != "0" {
		err = errors.New("短信发送失败" + string(body))
		return
	}

	if resp.StatusCode != http.StatusOK {
		err = errors.New("短信发送失败")
		return
	}
	aliSmsResult, err = srv.buildResult(string(body))
	return
}

func (srv *ChinaUnicomSms) buildResult(data string) (aliSmsResult *dysmsapi20170525.SendBatchSmsResponse, err error) {
	result, _ := url.ParseQuery(data)
	if _, ok := result["result"]; !ok {
		err = errors.New("短信发送失败")
		return
	}
	sendBatchSmsResponseBody := &dysmsapi20170525.SendBatchSmsResponseBody{}
	if result.Get("result") != "0" {
		sendBatchSmsResponseBody.SetCode(result.Get("result"))
	} else {
		sendBatchSmsResponseBody.SetCode("ok")
	}

	sendBatchSmsResponseBody.SetMessage(result.Get("description"))
	aliSmsResult = &dysmsapi20170525.SendBatchSmsResponse{}
	aliSmsResult.SetBody(sendBatchSmsResponseBody)
	return
}
