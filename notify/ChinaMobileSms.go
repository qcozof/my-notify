/**
* @Description:移动短信
e.g.
cmSms := ChinaMobileSms{
	TemplateId:"",
	Mobiles:"",
	Params:"",
}
cmSms.Send()
*/
package notify

import (
	"github.com/qcozof/my-notify/global"
	"github.com/qcozof/my-notify/utils"
	"github.com/qcozof/my-notify/utils/encrypt"

	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type ChinaMobileSms struct {
	TemplateId string `json:"templateId"` //模板ID
	Mobiles    string `json:"mobiles"`    //收信手机号码。英文逗号分隔，每批次限5000个号码，例：“13800138000,13800138001,13800138002”。
	Params     string `json:"params"`     //模板变量。格式：[“param1”,“param2”]，无变量模板填[""]。
}

type smsRequest struct {
	EcName     string `json:"ecName"`
	ApId       string `json:"ApId"`
	SecretKey  string `json:"secretKey"`
	TemplateId string `json:"templateId"`
	Mobiles    string `json:"mobiles"`
	Params     string `json:"params"`
	Sign       string `json:"sign"`
	AddSerial  string `json:"addSerial"`
	Mac        string `json:"mac"`
}

type smsResponse struct {
	Rspcod   string `json:"rspcod"`
	MgsGroup string `json:"mgsGroup"`
	Success  bool   `json:"success"`
}

func (srv *ChinaMobileSms) Send() error {
	//连接地址：https://****:****/sms/tmpsubmit
	//请求方式：post
	//数据类型：json（base64加密）
	var smsResponse smsResponse
	cfg := global.SERVER_CONFIG.ChinaMobileSmsConfig

	//模板为空时调用默认的注册模板
	if len(srv.TemplateId) == 0 {
		srv.TemplateId = cfg.TemplateId
	}

	b, err := json.Marshal(smsRequest{
		EcName:     cfg.EcName,
		ApId:       cfg.ApId,
		SecretKey:  cfg.SecretKey,
		TemplateId: srv.TemplateId,
		Mobiles:    srv.Mobiles,
		Params:     srv.Params,
		Sign:       cfg.Sign,
		AddSerial:  cfg.AddSerial,
		Mac:        srv.generateMac(),
	})

	if err != nil {
		return err
	}

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true} //忽略错误https证书
	b64Json := base64.StdEncoding.EncodeToString(b)
	wrapper := utils.PostJson(cfg.ApiUrl, b64Json, 30)
	body := wrapper.Body
	if wrapper.StatusCode != 200{
		return errors.New("短信发送失败："+ body)
	}

	err = json.Unmarshal([]byte(body), &smsResponse)
	if err != nil {
		return  errors.New(fmt.Sprintf("短信返回结果序列化失败：%s",err))
	}

	if !smsResponse.Success {
		return errors.New(fmt.Sprintf("短信发送失败 rspcod:%s mgsGroup:%s", smsResponse.Rspcod, smsResponse.MgsGroup))
	}
	return nil
}

func (srv *ChinaMobileSms) generateMac() string {
	//参数校验序列，生成方法：将ecName、ApId、secretKey、templateId、mobiles、params、sign、addSerial按序拼接（无间隔符），通过MD5（32位小写）计算出的值。
	cfg := global.SERVER_CONFIG.ChinaMobileSmsConfig
	str := fmt.Sprintf("%s%s%s%s%s%s%s%s", cfg.EcName, cfg.ApId, cfg.SecretKey, srv.TemplateId, srv.Mobiles, srv.Params, cfg.Sign, cfg.AddSerial)
	return encrypt.MD5V([]byte(str))
}
