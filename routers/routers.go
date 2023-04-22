package routers

import (
          "database/sql"
          
          "github.com/gorilla/mux"
          
          `StudentManagement/controllers`
          `StudentManagement/token`
)

func InitRouter(r *mux.Router, db *sql.DB, jwt *token.JWTMaker) {
          cfg := &controllers.Config{
                    DB:  db,
                    JWT: jwt,
          }
          InitStudentRouter(r, cfg)
          InitAuthRouter(r, cfg)
          InitCoursesRouter(r, cfg)
          
}
