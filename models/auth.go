package models

import (
          "database/sql"
          "errors"
          "fmt"
          "time"
          
          `StudentManagement/token`
)

var insertNewToken = "INSERT into token (id, username, role, issue_at, expired_at) VALUES (?, ?, ?, ?, ?)"
var loginUser = "SELECT login_user(?, ?)"
var changeUserPassword = `SELECT change_student_password(?, ?, ?) as shit`

type UserChangePassword struct {
          Username    string `json:"username" validate:"required"`
          Password    string `json:"password" validate:"required"`
          NewPassword string `json:"new_password" validate:"required"`
}

type UserLogin struct {
          Username string `json:"username" validate:"required"`
          Password string `json:"password" validate:"required"`
}

func (u *UserLogin) Login(jwt *token.JWTMaker, db *sql.DB) (map[string]interface{}, error) {
          fmt.Println("here is user", u)
          checkLoginStatus := "FAIL"
          err := db.QueryRow(loginUser, u.Username, u.Password).Scan(&checkLoginStatus)
          if err != nil {
                    fmt.Println("Error after login ", err)
                    return nil, err
          }
          if checkLoginStatus == "FAIL" {
                    return nil, errors.New("There is no user with provided credentials")
          }
          userRole := checkLoginStatus
          var d time.Duration = 10000000000000
          payload, token, err := jwt.CreateToken(u.Username, userRole, d)
          res := map[string]interface{}{
                    "payload": payload,
                    "token":   token,
          }
          db.QueryRow(insertNewToken, payload.ID, payload.Username, payload.Role, payload.IssuedAt, payload.ExpiredAt)
          return res, err
}

func Logout(jwt *token.JWTMaker, token string) (map[string]interface{}, error) {
          res, err := jwt.RevokeToken(token)
          return res, err
}

func (s *UserChangePassword) ChangePassword(db *sql.DB) error {
          rowAffected := 0
          err := db.QueryRow(changeUserPassword, s.Username, s.Password, s.NewPassword).Scan(&rowAffected)
          fmt.Println("function cal resssss", err, rowAffected)
          if err != nil {
                    return err
          }
          if rowAffected == 0 {
                    return errors.New("No user found with provided credentials!")
          }
          return err
}
