package messages

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func SendSms(msg string, phonenumber string) {

	path := fmt.Sprintf("%s?api_token=%s&from=%s&to=%s&body=%s",
		Sms_endpoint, Api_token, From, phonenumber, msg)

	res, err := http.Post(path, "application/json", nil)
	if err != nil {
		log.Printf("Error making request %v", err)
	}

	fmt.Println(res.Status)
}

func SendBulkSms(msg string, phonenumbers []string) {
	To := strings.Join(phonenumbers, ",")

	path := fmt.Sprintf("%s?api_token=%s&from=%s&to=%s&body=%s",
		Sms_endpoint, Api_token, From, To, msg)

	res, err := http.Post(path, "application/json", nil)
	if err != nil {
		log.Printf("Error making request %v", err)
	}
	fmt.Println(res.Status)
}
