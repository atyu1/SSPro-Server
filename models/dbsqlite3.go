package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

var db *gorm.DB

func InitDb(config DbAccess) {
	var err error
	db, err = gorm.Open("sqlite3", fmt.Sprintf("%s/%s", config.DbPath, config.DbName))
	if err != nil {
		log.Fatal("Connection to database failed!", fmt.Sprintf("%v", err))
	}

	//ToDo test user, remove it later!
	user := &User{Email:"test@test.com", Password:"test123#", Token:""}

	db.Debug().AutoMigrate(&Datapoint{})
	db.Debug().AutoMigrate(&User{})
	db.Debug().Create(&user)
}

func GetDb() *gorm.DB {
	return db
}
