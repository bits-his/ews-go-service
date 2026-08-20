package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/salemzii/cygio/app"
)

func main() {

	path, _ := os.Getwd()
	err := godotenv.Load(filepath.Join(path, ".env"))
	if err != nil {
		log.Println("Error loading .env file")
	}

	PRT := os.Getenv("PORT")
	router := gin.Default()

	router.GET("/welcome", welcome)
	router.POST("/createalert", app.ReceiveAlert)
	router.Run(fmt.Sprintf(":%s", PRT))
}

func welcome(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "Welcome to ews api",
	})
}
