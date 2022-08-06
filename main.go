package main

import (
	"github.com/gin-gonic/gin"
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

}

func welcome(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Hello welcome to Franka webhook",
	})
}
