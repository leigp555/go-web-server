package util

import (
	"github.com/jordan-wright/email"
	"net/smtp"
)

func SendEmail(emailNumber []string) (err error) {
	e := email.NewEmail()
	//设置发送方的邮箱
	e.From = "907090585@qq.com"
	// 设置接收方的邮箱
	e.To = emailNumber
	//设置主题
	e.Subject = "这是主题"
	//设置文件发送的内容
	e.HTML = []byte(`
    <h1><a href="http://www.topgoer.com/">go语言中文网站</a></h1>    
    `)
	//设置服务器相关的配置
	err = e.Send("smtp.qq.com:25", smtp.PlainAuth("", "907090585@qq.com", "cvxjaeubymkxbbic", "smtp.qq.com"))
	return
}
