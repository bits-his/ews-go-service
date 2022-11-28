package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/salemzii/cygio/facebook"
	"github.com/salemzii/cygio/messages"
	"github.com/salemzii/cygio/telegram"
	"github.com/streadway/amqp"
)

var (
	Wg     sync.WaitGroup
	Endpwg sync.WaitGroup
)

func ConsumeAlerts() {

	log.Println("Consumer Application")
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		log.Fatalf("Error connecting to rabbitMq %v", err)
	}
	defer conn.Close()

	fmt.Println("Successfully connected to rabbitmq")

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Error creating channel %v", err)
	}

	defer ch.Close()

	msgs, err := ch.Consume(
		"alert_queue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Error publishing queue %v", err)
	}
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Println(string(d.Body))

			var alert Alert
			if err := json.Unmarshal(d.Body, &alert); err != nil {
				log.Fatalf("%v", err)
			}
			CreateAlerts(alert)
			/*
				numbers := []string{}
				for _, v := range alert.Phones {
					numbers = append(numbers, v.Number)
				}
				messages.SendBulkSms(alert.Body, numbers)

				mails := []string{}

				for _, v := range alert.Mails {
					mails = append(mails, v.Address)
				}
				messages.SendMails(alert.Headline, alert.Body, mails)
			*/
		}
	}()

	<-forever

}

func SendNewsletterMail() {

	log.Println("Consumer Application")
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		log.Fatalf("Error connecting to rabbitMq %v", err)
	}
	defer conn.Close()

	fmt.Println("Successfully connected to rabbitmq")

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Error creating channel %v", err)
	}

	defer ch.Close()

	msgs, err := ch.Consume(
		"cipher",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Error publishing queue %v", err)
	}
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Println(string(d.Body))

			var subscriber struct {
				Email string `json:"email"`
			}

			if err := json.Unmarshal(d.Body, &subscriber); err != nil {
				log.Fatalf("%v", err)
			}

			messages.SendSSLMail("Drug Cipher Subscription",
				"Hi you have successfully subscribed to drug cipher's newsletter, you will receive subsequent update about our products",
				subscriber.Email)
		}
	}()

	<-forever
}

func ReceiveAlert(c *gin.Context) {

	var alert Alert
	if c.Request.Body != nil {
		c.BindJSON(&alert)
		log.Println(alert)
	}

	log.Println("Dispensing alerts to all listed platforms and urls")
	//defer CreateAlerts(alert)

	mails := []string{}

	for _, v := range alert.Mails {
		mails = append(mails, v.Address)
	}

	fmt.Println(mails)

	messages.SendMails(alert.Headline, alert.Body, mails)

	c.JSON(200, gin.H{
		"success": "Alerts distributed successfully",
	})
}

func CreateAlerts(alert Alert) {

	Platforms := alert.Platforms
	Urls := alert.Urls
	text := alert.Body

	log.Println(alert.Mails)
	mails := []string{}
	numbers := []string{}

	for _, v := range alert.Mails {
		mails = append(mails, v.Address)
	}

	for _, v := range alert.Phones {
		numbers = append(numbers, v.Number)
	}

	log.Println(mails)
	log.Println(numbers)

	messages.SendMails(alert.Headline, alert.Body, mails)
	messages.SendBulkSms(alert.Body, numbers)

	Wg.Add(len(Platforms))
	for _, v := range Platforms {
		switch v.Name {

		case "twitter":
			log.Println("Creating alert on twitter")
			/*
				go func(text string) {
					defer Wg.Done()
					twitter.CreateTweet(text)
				}(text)
			*/

		case "facebook":
			log.Println("Creating alert on facebook")
			go func(text string) {
				defer Wg.Done()
				facebook.PagePost(text)
			}(text)

		case "telegram":
			log.Println("Creating alert on telegram")
			go func(text string) {
				defer Wg.Done()
				_, err := telegram.SendChanMsg(text)
				if err != nil {
					log.Printf("Error publishing to telegram channel: %v", err)
				}
			}(text)

		case "whatsapp":
			log.Println("Creating alert on whatsapp")
			go func(text string) {
				defer Wg.Done()
				facebook.SendMsg()
			}(text)
		}
	}

	Wg.Wait()

	SendToEndpoints(Urls, text)

}

type AlertBody struct {
	Body string `json:"body"`
}

func SendToEndpoints(Urls []Url, content string) {
	alertBody := AlertBody{Body: content}
	data, err := json.Marshal(&alertBody)
	if err != nil {
		log.Println("Error encoding to json", err)
	}

	Endpwg.Add(len(Urls))

	for _, v := range Urls {
		go func(url Url) {
			defer Endpwg.Done()
			resp, err := http.Post(url.Uri, "application/json", bytes.NewBuffer(data))
			if err != nil {
				log.Printf("Error posting data to %s", err)
			}
			defer resp.Body.Close()
			log.Println(resp.Status)
		}(v)
	}

	Endpwg.Wait()
}
