package middleware

import (
          "context"
          "fmt"
          "log"
          "net/http"
          "strings"
          
          `StudentManagement/controllers`
          `StudentManagement/token`
)

type middleware func(http.HandlerFunc, *token.JWTMaker) http.HandlerFunc

func ChainMiddleware(f http.HandlerFunc, jwt *token.JWTMaker, m ...middleware) http.HandlerFunc {
          if len(m) == 0 {
                    return f
          }
          currentMiddleWare := m[0]
          return currentMiddleWare(ChainMiddleware(f, jwt, m[1:cap(m)]...), jwt)
}

func TokenMiddleware(next http.HandlerFunc, jwt *token.JWTMaker) http.HandlerFunc {
          return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                    log.Println(r.RequestURI)
                    authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
                    if len(authHeader) != 2 {
                              fmt.Println("Malformed token")
                              w.WriteHeader(http.StatusUnauthorized)
                              w.Write([]byte("Malformed Token"))
                              return
                    }
                    payload, err := jwt.VerifyToken(authHeader[1])
                    if err != nil {
                              controllers.RespondWithError(w, http.StatusUnauthorized, err.Error())
                              return
                    }
                    rcopy := r.WithContext(context.WithValue(r.Context(), "payload", payload))
                    next.ServeHTTP(w, rcopy)
          })
}

func ProfessorRoleMiddleware(next http.HandlerFunc, jwt *token.JWTMaker) http.HandlerFunc {
          return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                    log.Println(r.RequestURI)
                    user := r.Context().Value("payload")
                    token := user.(*token.Payload)
                    role := token.Role
                    fmt.Println("role is", role)
                    if role != "PROFESSOR" {
                              controllers.RespondWithError(w, http.StatusInternalServerError, "Role not granted to you!")
                              return
                    }
                    next.ServeHTTP(w, r)
          })
}
