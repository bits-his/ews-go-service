package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/salemzii/cygio/app"
	"github.com/streadway/amqp"
)

/*
type Alert struct {
	Topic   string `json:"topic"`
	Message string `json:"message"`
}
*/

func main() {

	ConnectRMQ()
}

func ConnectRMQ() {

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
			var alert app.Alert
			if err := json.Unmarshal(d.Body, &alert); err != nil {
				log.Fatalf("%v", err)
			}
			app.CreateAlerts(alert)
		}
	}()

	<-forever

}

func welcome(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Hello welcome to Franka webhook",
	})
}

/*
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {



		router := gin.Default()

		router.GET("/", welcome)
		router.POST("/createalert", app.ReceiveAlert)

		router.Run()


			t1 := time.Now()
			messages.SendMails("Testing mails; if mails still work :(", "Hello everyone @ brainstorm",
				[]string{"salemododa2@gmail.com", "harunakadiri702@gmail.com",
					"robtyler0701@gmail.com", "davidbill0701@gmail.com",
					"jacobmamudu044@gmail.com", "chopfastfast@gmail.com", "issatoyin@gmail.com", "franksultan48@gmail.com"})
			t2 := time.Now()
			log.Printf("Total time taken %v", t2.Sub(t1))


	//messages.SendMails("Testing mails; if mails still work :(", "Hello everyone @ brainstorm", []string{"salemododa2@gmail.com", "robtyler0701@gmail.com"})
}
*/
