package routers

import (
          "fmt"
          "net/http"
          
          "github.com/gorilla/mux"
          
          `StudentManagement/controllers`
          `StudentManagement/middleware`
)

func InitCoursesRouter(r *mux.Router, cfg *controllers.Config) {
          fmt.Println("Initialize courses route...")
          r.PathPrefix("/courses").Subrouter().Handle("/", middleware.TokenMiddleware(http.HandlerFunc(cfg.GetCoursesHandler), cfg.JWT)).Methods("GET")
          r.PathPrefix("/courses").Subrouter().Handle("/{course_id}/students/", middleware.ChainMiddleware(
                    http.HandlerFunc(cfg.GetCourseStudentHandler),
                    cfg.JWT,
                    middleware.TokenMiddleware,
                    middleware.ProfessorRoleMiddleware,
          )).Methods("GET")
          r.PathPrefix("/courses").Subrouter().Handle("/{course_id}/exam_hw/", middleware.ChainMiddleware(
                    http.HandlerFunc(cfg.GetCourseHWExamHandler),
                    cfg.JWT,
                    middleware.TokenMiddleware,
                    middleware.ProfessorRoleMiddleware,
          )).Methods("GET")
          r.PathPrefix("/courses").Subrouter().Handle("/exam/", middleware.ChainMiddleware(
                    http.HandlerFunc(cfg.CreateCourseExamHandler),
                    cfg.JWT,
                    middleware.TokenMiddleware,
                    middleware.ProfessorRoleMiddleware,
          )).Methods("POST")
          r.PathPrefix("/courses").Subrouter().Handle("/exam/question/", middleware.ChainMiddleware(
                    http.HandlerFunc(cfg.AddExamQuestionHandler),
                    cfg.JWT,
                    middleware.TokenMiddleware,
                    middleware.ProfessorRoleMiddleware,
          )).Methods("POST")
          r.PathPrefix("/courses").Subrouter().Handle("/exam/question/submit/", middleware.ChainMiddleware(
                    http.HandlerFunc(cfg.SubmitExamAnswerHandler),
                    cfg.JWT,
                    middleware.TokenMiddleware,
          )).Methods("POST")
          r.PathPrefix("/courses").Subrouter().Handle("/exam/{exam_id}/score/", middleware.ChainMiddleware(
                    http.HandlerFunc(cfg.GetExamScoreHandler),
                    cfg.JWT,
                    middleware.TokenMiddleware,
          )).Methods("GET")
}
