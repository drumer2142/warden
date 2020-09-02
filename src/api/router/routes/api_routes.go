package routes

import (
  "net/http"
  "github.com/drumer2142/warden/src/api/controllers"
)

var api_routes = []Route{
  Route{
    URI: "/sign-in",
    Method: http.MethodPost,
    Controller: controllers.SignIn,
  },
  Route{
    URI: "/auth-route",
    Method: http.MethodPost,
    Controller: controllers.AuthRoute,
  },
  Route{
    URI: "/refresh",
    Method: http.MethodPost,
    Controller: controllers.RefreshToken,
  },
}
