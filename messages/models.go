package messages

import "os"

var (
	Sms_endpoint  = "https://www.bulksmsnigeria.com/api/v1/sms/create"
	Api_token     = os.Getenv("SMS_TOKEN")
	From          = "EWS"
	Body          = ""
	Gateway       = "0"
	Append_sender = "EWS"
)
