package api

import (
  "fmt"
  "log"
  "net/http"
  "github.com/drumer2142/warden/src/config"
  "github.com/drumer2142/warden/src/api/router"
)

func init(){
  config.Load()
}

func main(){
  port := config.APP_PORT
  router := router.New()

  fmt.Printf("\nListening [::]:%d \n", port)
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
