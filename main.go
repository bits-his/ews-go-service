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
		log.Fatal("Error loading .env file")
	}

	PRT := os.Getenv("PRT")
	router := gin.Default()

	router.POST("/createalert", app.ReceiveAlert)
	router.Run(fmt.Sprintf(":%s", PRT))
}
