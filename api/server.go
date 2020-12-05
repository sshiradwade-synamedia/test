package api

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sanjayshr/login/api/controllers"
)

var server = controllers.Server{}


func Run() {

  var err error
  err = godotenv.Load()
  if err != nil {
    log.Fatal("Error getting environment, %v", err)
  }

  server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))


	server.Run(":8080")
}
