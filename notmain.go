package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/salemzii/cygio/app"
	"github.com/salemzii/cygio/messages"
	"github.com/streadway/amqp"
)

/*
type Alert struct {
	Topic   string `json:"topic"`
	Message string `json:"message"`
}
*/

func main() {
	ConnectRMQ2()
}

func ConnectRMQ2() {

	log.Println("Consumer Application")
	conn, err := amqp.Dial("amqps://uobviquo:nhFtcBQeVLiV9HKN1kfzQzKgW1brulyu@beaver.rmq.cloudamqp.com/uobviquo")
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

			var alert app.Alert
			if err := json.Unmarshal(d.Body, &alert); err != nil {
				log.Fatalf("%v", err)
			}
			//app.CreateAlerts(alert)
			/*
				numbers := []string{}
				for _, v := range alert.Phones {
					numbers = append(numbers, v.Number)
				}
				messages.SendBulkSms(alert.Body, numbers)
			*/
			mails := []string{}

			for _, v := range alert.Mails {
				mails = append(mails, v.Address)
			}
			messages.SendMails(alert.Headline, alert.Body, mails)

			numbers := []string{}
			for _, v := range alert.Phones {
				numbers = append(numbers, v.Number)
			}
			messages.SendBulkSms(alert.Body, numbers[0:2])
		}
	}()

	<-forever

}
