package controllers

import (
  "time"
  "net/http"
  "encoding/json"
  "github.com/dgrijalva/jwt-go"
  "github.com/drumer2142/warden/src/api/models"
  "github.com/drumer2142/warden/src/config"
)

var jwtKey = config.JWTKEY

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

func SignIn(w http.ResponseWriter, r *http.Request){
  var creds models.Credentials

  err := json.NewDecoder(r.Body).Decode(&creds)
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    return
  }

  expectedPassword, ok := users[creds.Username]

  if !ok || expectedPassword != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

  expirationTime := time.Now().Add(5 * time.Minute)
  claims := &models.Claims{
    Username: creds.Username,
    StandardClaims: jwt.StandardClaims{
      ExpiresAt: expirationTime.Unix(),
    },    
  }

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  tokenString, err := token.SignedString(jwtKey)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    return
  }
  
  http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}
