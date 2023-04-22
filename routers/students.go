package routers

import (
          "fmt"
          "net/http"
          
          "github.com/gorilla/mux"
          
          `StudentManagement/controllers`
          `StudentManagement/middleware`
)

func InitStudentRouter(r *mux.Router, cfg *controllers.Config) {
          fmt.Println("Initialize student route...")
          r.PathPrefix("/student").Subrouter().Handle("/test", middleware.TokenMiddleware(http.HandlerFunc(cfg.TestHandler), cfg.JWT)).Methods("GET")
}
