package models

import (
  "github.com/dgrijalva/jwt-go"
  "time"
)

type Credentials struct {
  ID       uint64 `gorm:"primary_id;auto_increment" json:"id"`
  Password string `gorm:"column:password;not null;type:varchar(255)" json:"password"`
  Username string `gorm:"column:username;not null;type:varchar(255)" json:"username"`
}

type Claims struct {
  Username string `json:"username"`
  jwt.StandardClaims
}

type AuthToken struct {
  Name string `json:"name"`
  Value string `json:"value"`
  Expires time.Time `json:"expires"`
}