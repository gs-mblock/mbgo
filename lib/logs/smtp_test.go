// Copyright 2014 beego Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logs

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"
	"testing"
	//"time"
)

// func TestSmtp(t *testing.T) {
// 	log := NewLogger(10000)
// 	log.SetLogger("smtp", `{"username":"beegotest@gmail.com","password":"xxxxxxxx","host":"smtp.gmail.com:587","sendTos":["xiemengjun@gmail.com"]}`)
// 	log.Critical("sendmail critical")
// 	time.Sleep(time.Second * 30)
// }

func TestSmtp2(t *testing.T) {
	// Choose auth method and set it up
	auth := smtp.PlainAuth("", "piotr@mailtrap.io", "extremely_secret_pass", "smtp.mailtrap.io")

	// Here we do it all: connect to our server, set up a message and send it
	to := []string{"56487685@qq.com"}
	msg := []byte("To: 56487685@qq.com\r\n" +
		"Subject: Why are you not using Mailtrap yet?\r\n" +
		"\r\n" +
		"Here’s the space for our great sales pitch\r\n")
	err := smtp.SendMail("smtp.mailtrap.io:25", auth, "piotr@mailtrap.io", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}

func TestSmtp_qq(t *testing.T) {
	setMail()
}

func setMail() {
	// 邮箱地址
	UserEmail := "37907909@qq.com"
	// 端口号，:25也行
	MailSMTPPort := ":587"
	//邮箱的授权码，去邮箱自己获取
	MailPassword := "axs89864jdqhpfjzmzywbhda"
	// 此处填写SMTP服务器
	MailSMTPHost := "smtp.qq.com"
	auth := smtp.PlainAuth("", UserEmail, MailPassword, MailSMTPHost)
	to := []string{"liam.jiang@makeblock.com"}
	nickname := "发送人名称"
	user := UserEmail

	subject := "我是我相信"
	contentType := "Content-Type: text/plain; charset=UTF-8"

	body := "测试一下"
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
		"<" + user + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	err := smtp.SendMail(MailSMTPHost+MailSMTPPort, auth, user, to, msg)
	if err != nil {
		fmt.Printf("send mail error: %v", err)
	}
}
