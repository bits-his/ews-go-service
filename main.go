package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/salemzii/cygio/messages"
)

func main() {

	/*
		router := gin.Default()

		router.GET("/", welcome)
		router.POST("/createalert", app.ReceiveAlert)

		router.Run()
		[]string{"2347014327332", "2348053503763", "2347062942291"}
		"2347014327332", "2348053503763","2347062942291"
	*/
	t1 := time.Now()
	messages.SendMails("Hello everyone @ brainstorm",
		[]string{"issatoyin@gmail.com", "salemododa2@gmail.com", "harunakadiri702@gmail.com",
			"robtyler0701@gmail.com", "davidbill0701@gmail.com"})
	t2 := time.Now()
	log.Printf("Total time taken %v", t2.Sub(t1))
}

func welcome(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Hello welcome to Franka webhook",
	})
}
