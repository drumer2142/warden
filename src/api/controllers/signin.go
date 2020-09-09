package controllers

import (
  "time"
  "net/http"
  "encoding/json"
  "github.com/dgrijalva/jwt-go"
  "github.com/drumer2142/warden/src/api/models"
  "github.com/drumer2142/warden/src/config"
  "github.com/drumer2142/warden/src/api/database"
  "github.com/drumer2142/warden/src/api/handler"
)

var jwtKey = config.JWTKEY

func SignIn(w http.ResponseWriter, r *http.Request){
  var creds models.Credentials

  // decode incoming json credentials on signin
  err := json.NewDecoder(r.Body).Decode(&creds)
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    return
  }

  // connect to the database
  db, err := database.Connect()
  if err != nil {
		handler.ResponseJSON(w, http.StatusInternalServerError, err)
		return
  }
  result := db.Debug().Where("username = ? AND password = ?", creds.Username,creds.Password).First(&creds).RecordNotFound()

  if result == true {
    handler.ResponseJSON(w, http.StatusUnauthorized, result)
    return
  }

  // create the token timeout and claims
  expirationTime := time.Now().Add(5 * time.Minute)
  claims := &models.Claims{
    Username: creds.Username,
    StandardClaims: jwt.StandardClaims{
      ExpiresAt: expirationTime.Unix(),
    },    
  }

  // create the token
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  tokenString, err := token.SignedString(jwtKey)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    return
  }
  
  // set a cookie if used that way else return a json response
  // http.SetCookie(w, &http.Cookie{ 
  //   Name:    "token",
	// 	Value:   tokenString,
	// 	Expires: expirationTime,
  // })

  cookie_return := models.AuthToken{
   Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
  }

  handler.ResponseJSON(w, http.StatusOK, cookie_return)
}
