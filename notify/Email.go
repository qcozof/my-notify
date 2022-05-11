/**
* @Description:邮件发送

/*
例：
    var dd = &notify.Email{
	MailTo: []string{"ww@qq.cn",},
	Subject: "发件测试",
	Body: "发件内容",
}

res,err := dd.Send()
fmt.Println(res,err)
*/
package notify


import (
	"errors"
	"github.com/qcozof/my-notify/global"
	"gopkg.in/gomail.v2"
)

type Email struct {
	MailTo  []string //接收邮件地址
	Subject string   //标题
	Body    string   //内容
	host     string //主机
	port     int    //端口
	username string
	password string
}

func (srv *Email) Send()  (res string,err error) {

/*	srv.host="smtp.exmail.qq.com"
	srv.port="465"
	srv.username = "xxx@qq.cn"
	srv.password = "123456"*/
	config := global.SERVER_CONFIG.EmailConfig
	srv.host = config.Host
	srv.port = config.Port
	srv.username = config.Username
	srv.password = config.Password

	if len(srv.host) == 0 || len(srv.username) == 0 || len(srv.password) == 0 {
		return "",errors.New("host,port,username,password不能为空。")
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "<"+srv.username+">")
	m.SetHeader("To", srv.MailTo...)
	m.SetHeader("Subject", srv.Subject)
	m.SetBody("text/html", srv.Body)
	d := gomail.NewDialer(srv.host, srv.port, srv.username, srv.password)
	err = d.DialAndSend(m)

	if err ==nil{
		return "发送成功",nil
	}
	return "发送失败",err
}