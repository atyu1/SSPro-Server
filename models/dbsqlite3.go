package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang/glog"
)

var db *gorm.DB

func InitDb(config DbAccess) {
	var err error
	db, err = gorm.Open("sqlite3", fmt.Sprintf("%s/%s", config.DbPath, config.DbName))
	if err != nil {
		glog.Fatalf("Connection to database failed! %v", err)
	}

	//ToDo test user, remove it later!
	user := &User{Email: "test@test.com", Password: "test123#", Token: ""}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	glog.Infof("ToDo, remove this part later! Default user created: %v", user)

	db.Debug().AutoMigrate(&Datapoint{})
	db.Debug().AutoMigrate(&User{})
	db.Debug().Create(&user)
}

func GetDb() *gorm.DB {
	return db
}
