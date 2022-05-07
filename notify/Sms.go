/**
* @Description:阿里云发短信
 */
package notify


/*
	var dd = notify.AliSms{
		PhoneNumberJson: "[\"1333333333\"]",
		TemplateParamJson: "",
	}
	res,err := dd.Send()
	fmt.Println(res,err)

*/

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/qcozof/my-notify/global"
)

//https://help.aliyun.com/document_detail/102364.html?spm=a2c4g.11186623.6.629.205e1e08YS98l3

type AliSms struct {
	//RegionId        string //区域ID 如 "cn-qingdao"
	accessKeyId     string
	accessKeySecret string

	PhoneNumberJson     string //电话号码 ["1590000****","1350000****"]
	SignNameJson        string //短信签名名称，JSON数组格式。
	TemplateCode        string //短信模板CODE。如 SMS_152550005
	TemplateParamJson   string //短信模板变量对应的实际值，JSON格式。 [{"name":"TemplateParamJson"},{"name":"TemplateParamJson"}]
	SmsUpExtendCodeJson string //上行短信扩展码，JSON数组格式。无特殊需要此字段的用户请忽略此字段。
}

//发送短信, templateCode模板编码(默认调用配置中的)
func  (srv *AliSms)Send (templateCode string) (aliSmsResult *dysmsapi20170525.SendBatchSmsResponse,err error) {
	config := global.SERVER_CONFIG.SmsConfig
	cfg := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: &config.AccessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: &config.AccessKeySecret,
	}
	// 访问的域名
	cfg.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	client := &dysmsapi20170525.Client{}
	client, err = dysmsapi20170525.NewClient(cfg)
	//return client, err


	//client, err := srv.CreateClient(tea.String("accessKeyId"), tea.String("accessKeySecret"))
	if err != nil {
		return
	}

	//模板为空时调用默认的注册模板
	if len(templateCode) == 0 {
		templateCode = config.TemplateCode
	}

	sendBatchSmsRequest := &dysmsapi20170525.SendBatchSmsRequest{
		PhoneNumberJson: &srv.PhoneNumberJson,
		SignNameJson: &config.SignNameJson,
		TemplateCode: &templateCode,
		TemplateParamJson: &srv.TemplateParamJson,
		//SmsUpExtendCodeJson: &srv.SmsUpExtendCodeJson,
	}
	// 复制代码运行请自行打印 API 的返回值
	aliSmsResult, err = client.SendBatchSms(sendBatchSmsRequest)

	return
}
