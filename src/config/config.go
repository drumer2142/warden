package config

import (
  "os"
  "log"
  "fmt"
  "strconv"
  "github.com/joho/godotenv"
)

var (
  APPPORT = 0
  DBDRIVER = ""
  DBURL = ""
  JWTKEY = []byte("")
)

func Load(){
  var err error
  err = godotenv.Load(".env")
  if err != nil {
    log.Fatal(err)
  }
  APPPORT, err = strconv.Atoi(os.Getenv("APP_PORT"))
  if err != nil {
    APPPORT = 8888
  }
  DBDRIVER = os.Getenv("DB_DRIVER")
  DBURL = fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
      os.Getenv("DB_USER"),
      os.Getenv("DB_PASS"),
      os.Getenv("DB_HOST"),
      os.Getenv("DB_PORT"),
      os.Getenv("DB_NAME"),
  )
  JWTKEY = []byte(os.Getenv("JWTKEY"))
}
