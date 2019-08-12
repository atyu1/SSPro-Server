package datapoints

import (
	"fmt"
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func InitDb(config DbAccess) {
	db, err := gorm.Open("sqlite3", fmt.Sprintf("%s/%s", config.DbPath, config.DbName))
	if err != nil {
		log.Fatal("Connection to database failed!")
	}

	db.Debug().AutoMigrate(&Datapoint{})
}

func GetDb() *gorm.DB {
	return db
}
