package main

import (
	"github.com/gin-gonic/gin"
	"github.com/salemzii/cygio/app"
)

func main() {

	router := gin.Default()

	router.GET("/", welcome)
	router.POST("/createalert", app.ReceiveAlert)

	router.Run()

	/*
		t1 := time.Now()
		messages.SendMails("Testing mails; if mails still work :(", "Hello everyone @ brainstorm",
			[]string{"salemododa2@gmail.com", "harunakadiri702@gmail.com",
				"robtyler0701@gmail.com", "davidbill0701@gmail.com",
				"jacobmamudu044@gmail.com", "chopfastfast@gmail.com", "issatoyin@gmail.com", "franksultan48@gmail.com"})
		t2 := time.Now()
		log.Printf("Total time taken %v", t2.Sub(t1))
	*/

	//messages.SendMails("Testing mails; if mails still work :(", "Hello everyone @ brainstorm", []string{"salemododa2@gmail.com", "robtyler0701@gmail.com"})
}

func welcome(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Hello welcome to Franka webhook",
	})
}
