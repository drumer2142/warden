package migrations

import (
  "log"
  "github.com/drumer2142/warden/src/api/database"
  "github.com/drumer2142/warden/src/api/models"
)


var err error

func Load() {
  db, err := database.Connect()
	if err != nil {
    log.Fatalf("\nFailed to connect to db....\n", err.Error())
  }
  defer db.Close()

  // Drop table if exits
  err = db.Debug().DropTableIfExists(&models.Credentials{}).Error
  if err != nil {
    log.Fatal(err)
  }

  // Migrate the schema
  err = db.Debug().AutoMigrate(&models.Credentials{}).Error
  if err != nil {
    log.Fatal(err)
  }

  // insert demo values to the db
  for i, _ := range creds {

    err = db.Debug().Model(&models.Credentials{}).Create(&creds[i]).Error
    if err != nil {
      log.Fatal(err)
    }

  }

}
