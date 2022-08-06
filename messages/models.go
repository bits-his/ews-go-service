package messages

import "os"

var (
	Sms_endpoint  = "https://www.bulksmsnigeria.com/api/v1/sms/create"
	From          = "EWS"
	Gateway       = "0"
	Append_sender = "EWS"
	Api_token     = os.Getenv("SMS_TOKEN")
	From_mail     = os.Getenv("FROM_MAIL")
	Mail_password = os.Getenv("MAIL_PASSWORD")
	SMTP_Host     = os.Getenv("HOST")
)
