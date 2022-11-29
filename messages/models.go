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
	Append_sender = "EWS ALERT DISPATCHER"
	Api_token     = os.Getenv("SMS_TOKEN")

	From_mail     = os.Getenv("EWS_MAIL")
	Mail_password = os.Getenv("EWS_MAIL_PASSWORD")
	SMTP_Host     = os.Getenv("EWS_MAIL_HOST")

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

func deprecatedInit() {
	from = &mail.Address{Name: "Ews mail", Address: From_mail}
	auth = smtp.PlainAuth("", From_mail, Mail_password, SMTP_Host)
	tlsconfig = &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         SMTP_Host,
	}
}

func NewContainer() *Container {
	return &Container{
		Headers: make(map[string]string),
	}
}
