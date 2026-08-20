package messages

import (
	"fmt"
	"log"

	twilio "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

/*
	export TWILIO_ACCOUNT_SID=xxxxxxxxx
	export TWILIO_AUTH_TOKEN=xxxxxxxxx
	export TWILIO_PHONE_NUMBER=xxxxxxxxx
	export TO_PHONE_NUMBER=xxxxxxxxx
*/

var (
	TWILIO_ACCOUNT_SID  = "AC731b80dc72f3f941010104afd18367d3"
	TWILIO_AUTH_TOKEN   = "tEHND17xUCZx1aHdXjXdIXG3ODpARBdOw27xiUlM"
	TWILIO_PHONE_NUMBER = "+13854755581"
	TO_PHONE_NUMBER     = "xxxxxxxxx"
)

func SendSmsTwilio() {
	client := twilio.NewRestClient()
	log.Println(client.Client.AccountSid())

	params := &openapi.CreateMessageParams{}
	//params.SetTo(os.Getenv("TO_PHONE_NUMBER"))
	params.SetTo("+2347014327332")
	params.SetFrom(TWILIO_PHONE_NUMBER)
	params.SetBody("Hello world, testing ews with golang")

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("SMS sent successfully!")
	}
}
