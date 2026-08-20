package telegram

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	CHANNEL_NAME   = "-1001847384984"
	TELEGRAM_TOKEN = os.Getenv("TELEGRAM_TOKEN")
	MESSAGE        = "Hello World"
	ENDPOINT       = "https://api.telegram.org/bot"
)

type SendChaMsgResp struct {
	Ok     bool   `json:"ok"`
	Result Result `json:"result"`
}

type Result struct {
	MessageId  int    `json:"message_id"`
	SenderChat Chat   `json:"sender_chat"`
	Chat       Chat   `json:"chat"`
	Date       int    `json:"date"`
	Text       string `json:"text"`
}
type Chat struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Type  string `json:"type"`
}

/*
	In case you need to private the telegram group, you can first make the channel public, then
	publish to the public group; the response payload will contain the channel_id of form "-1001846384985" which you can
	use to send message to the private channel
*/

func HelloWorld() {
	url := fmt.Sprintf(ENDPOINT+"%s/sendMessage?chat_id=%s&text=%s",
		TELEGRAM_TOKEN, CHANNEL_NAME, MESSAGE)

	log.Println(url)
	resp, err := http.Get(url)

	if err != nil {
		log.Fatalf("Error %v", err)
	}

	log.Println(resp.Status)
	log.Println(resp.Body)
}

func SendChanMsg(text string) (ok bool, err error) {
	url := fmt.Sprintf(ENDPOINT+"%s/sendMessage?chat_id=%s&text=%s",
		TELEGRAM_TOKEN, CHANNEL_NAME, text)

	log.Println(url)

	var sendChaMsgResp SendChaMsgResp

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error %v", err)
	}

	// decode the response body to a  bytes
	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(respByte, &sendChaMsgResp); err != nil {
		log.Fatalf("Error %v", err)
		return false, err
	}
	log.Println(sendChaMsgResp)
	return sendChaMsgResp.Ok, nil
}
