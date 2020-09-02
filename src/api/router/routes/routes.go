package routes

import (
  "net/http"
  "github.com/gorilla/mux"
  // if middleware ...
)

type Route struct {
  URI           string
  Method        string
  Controller    func(w http.ResponseWriter, r *http.Request)
}

func LoadRoutes() []Route {
  routes := api_routes
  return routes
}

func SetupRoutes(r *mux.Router) *mux.Router{
  for _, route := range LoadRoutes() {
    r.HandleFunc(route.URI, route.Controller).Methods(route.Method)
  }
  return r
}
