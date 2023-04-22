package controllers

import (
          "fmt"
          "net/http"
          
          "github.com/go-playground/validator/v10"
          
          `StudentManagement/token`
)

var Validator = validator.New()

type testStruct struct {
          Test string `json:"test"`
}

func (cfg *Config) TestHandler(w http.ResponseWriter, r *http.Request) {
          
          user := r.Context().Value("payload")
          fmt.Println("==========")
          token := user.(*token.Payload)
          username := token.Username
          fmt.Println(username)
          
          RespondWithJSON(w, http.StatusOK, testStruct{
                    Test: "hello world",
          })
}
