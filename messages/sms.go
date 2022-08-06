package messages

import (
	"fmt"
	"log"
	"net/http"
)

func SendSms(msg string, phonenumbers ...string) {

	lenghtSmsList := len(phonenumbers)
	To := ""

	for i, v := range phonenumbers {

		if i == lenghtSmsList-1 {
			To += v
			break
		}
		To += fmt.Sprintf("%s,", v)
	}

	path := fmt.Sprintf("%s?api_token=%s&from=%s&to=%s&body=%s",
		Sms_endpoint, Api_token, From, To, msg)

	res, err := http.Post(path, "application/json", nil)
	if err != nil {
		log.Printf("Error making request %v", err)
	}
	fmt.Println(res.Status)
}
