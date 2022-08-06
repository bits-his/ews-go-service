package messages

import (
	"crypto/tls"
	"net/mail"
	"net/smtp"
	"os"
	"sync"
)

var (
	Sms_endpoint  = "https://www.bulksmsnigeria.com/api/v1/sms/create"
	From          = "EWS"
	Gateway       = "0"
	Append_sender = "EWS"
	Api_token     = os.Getenv("SMS_TOKEN")
	From_mail     = os.Getenv("FROM_MAIL")
	Mail_password = os.Getenv("MAIL_PASSWORD")
	SMTP_Host     = os.Getenv("HOST")

	ErrSendingSmsRequest = "Encountered Error While Sending Sms Request"
	ErrSendingMail       = "Encountered Error While Sending Mail"

	Mail_subject string
	Mail_body    string
	mailwg       sync.WaitGroup
)

type Container struct {
	m       sync.Mutex
	Headers map[string]string
}

var (
	from      *mail.Address
	auth      smtp.Auth
	tlsconfig *tls.Config
)

func init() {
	from = &mail.Address{Name: "EWS", Address: From_mail}
	auth = smtp.PlainAuth("", From_mail, Mail_password, SMTP_Host)
	tlsconfig = &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         SMTP_Host,
	}
}
