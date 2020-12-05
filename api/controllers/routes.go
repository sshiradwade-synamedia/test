package controllers

import "github.com/sanjayshr/login/api/middlewares"


func (s *Server) initializeRoutes() {

  // User routes
  s.Router.HandleFunc("/register", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
}
