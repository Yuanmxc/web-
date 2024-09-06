package gmail

import (
	"bytes"
	"fmt"
	"github.com/jordan-wright/email"
	"log"
	"math/rand"
	"net/smtp"
	"strconv"
)

var Mail *email.Email

// New 初始化用户名和密码
func New() {
	Mail = email.NewEmail()
	Mail.From = "kangningwang254@gmail.com"
	Mail.Subject = "TTMS"
}

// Send 标题 文本 目标邮箱
func Send(toId string, verification string) {
	// 设置接收方的邮箱
	Mail.To = []string{toId}
	//设置主题
	Mail.Subject = "TTMS"
	//设置文件发送的内容
	text := fmt.Sprintf("[TTMS] 验证码：%s,5分钟后过期，请勿转发", verification)
	Mail.Text = []byte(text)
	//设置服务器相关的配置
	log.Println("start send")
	err := Mail.Send("smtp.gmail.com:587", smtp.PlainAuth("", "kangningwang254@gmail.com", "wdbouxvvlwhcqvgg", "smtp.gmail.com"))
	for err != nil {
		log.Println(err)
		log.Println("验证码没发过去，我重试一次")
		err = Mail.Send("smtp.gmail.com:587", smtp.PlainAuth("", "kangningwang254@gmail.com", "wdbouxvvlwhcqvgg", "smtp.gmail.com"))
	}
	log.Println("send complete")
}
func GetVerification() string {
	ans := bytes.Buffer{}
	for i := 0; i < 6; i++ {
		ans.WriteString(strconv.Itoa(rand.Intn(10)))
	}
	return ans.String()
}
