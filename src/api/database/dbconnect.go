package database

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "github.com/drumer2142/warden/src/config"
)

func Connect() (*gorm.DB, error){
  db, err := gorm.Open(config.DBDRIVER, config.DBURL)
  if err != nil {
    return nil, err
  }
  return db, nil
}
