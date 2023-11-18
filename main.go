package main

import (
	"log"
	"os"
	"rest-2/app/router"
	"rest-2/config"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	config.InitLog()
}

func main() {
	log.Output(1, "hello")
	port := os.Getenv("PORT")

	init := config.Init()
	app := router.Init(init)

	app.Run(":" + port)
}
