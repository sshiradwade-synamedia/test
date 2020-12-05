package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sanjayshr/login/api/models"
	"github.com/sanjayshr/login/api/responses"
	"github.com/sanjayshr/login/api/utils/formaterror"
)

func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {

  body, err := ioutil.ReadAll(r.Body)
  if err != nil {
      responses.ERROR(w, http.StatusUnprocessableEntity, err)    
  }

  user := models.User{}
  
  err = json.Unmarshal(body, &user)

  if err != nil {
    responses.ERROR(w, http.StatusUnprocessableEntity, err)
    return
  
  user.Prepare()

  userCreated, err := user.SaveUser(server.DB)
  if err != nil {
    formattedError := formaterror.FormatError(err.Error())

    responses.ERROR(w, http.StatusInternalServerError, formattedError)
    return
  }

  w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.RequestURI, userCreated))
  responses.JSON(w, http.StatusCreated, userCreated) 
}
}
