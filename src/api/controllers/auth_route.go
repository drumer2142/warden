package controllers

import (
	_"log"
	"net/http"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/drumer2142/warden/src/api/models"
	"github.com/drumer2142/warden/src/api/handler"
)

func AuthRoute(w http.ResponseWriter, r *http.Request){
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

	w.WriteHeader(http.StatusOK)
	return
}
