package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/sanjayshr/login/api/models"

	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
)


type Server struct {
  DB *gorm.DB
  Router *mux.Router
}

// Initialize -
func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
  
  var err error
 
  DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
   fmt.Println(Dbdriver) 
  server.DB, err = gorm.Open(Dbdriver, DBURL)
  if err != nil {
    fmt.Printf("Cannot conect to %s database", Dbdriver)
    log.Fatal("This is the error:", err)
  } else {
       fmt.Printf("Database connection successful")
  }

  server.DB.Debug().AutoMigrate(&models.User{})

  server.Router = mux.NewRouter()

  server.initializeRoutes()
}


func (server *Server) Run(addr string){
  fmt.Println("Listening to port 8080")
  log.Fatal(http.ListenAndServe(addr, server.Router))
}
