package router

import (
  "github.com/drumer2142/warden/src/api/router/routes"
	"github.com/gorilla/mux"
)

func New() *mux.Router{
  r := mux.NewRouter().StrictSlash(true)
  return routes.SetupRoutes(r)
}
