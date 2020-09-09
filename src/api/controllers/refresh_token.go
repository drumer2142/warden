package controllers

import (
	"log"
	"time"
	"net/http"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/drumer2142/warden/src/api/models"
	"github.com/drumer2142/warden/src/api/handler"
 )

func RefreshToken(w http.ResponseWriter, r *http.Request){

	var authtkn models.AuthToken

	err := json.NewDecoder(r.Body).Decode(&authtkn)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if authtkn.Name != "token"{
		handler.ResponseError(w, http.StatusBadRequest, "Bad Token Given")
		return
	}

	tknStr := authtkn.Value
	log.Println(tknStr)

	claims := &models.Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Now, create a new token for the current use, with a renewed expiration time
	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	cookie_return := models.AuthToken{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	}

	handler.ResponseJSON(w, http.StatusOK, cookie_return)
}