package models

import (
	"github.com/atyu1/SSPro-Server/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang/glog"
)

//ToDo Add to config file
var TokenPassword string = "ThisIsTokenPasswordTemporary"

// JWT User struct
type Token struct {
	UserId uint
	jwt.StandardClaims
}

// User is a struct to define User
type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token";sql:"-"`
}

func Login(email, password string) map[string]interface{} {
	user := &User{}

	err := GetDb().Table("users").Where("email = ?", email).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.Message(false, "Email address not found")
			glog.Error("Email address not found")
		}
		glog.Errorf("Database connection issue: %v", err)
		return utils.Message(false, "Database connection error")
	}

	glog.Infof("Retrieved user for email(%s): %v)",email, user)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		glog.Errorf("Password is not matching, %v", err)
		return utils.Message(false, "Invalid login credentials")
	}

	user.Password = "" //Remove password from memory

	token := &Token{UserId: user.ID}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), token)
	tokenString, _ := jwtToken.SignedString([]byte(TokenPassword))
	user.Token = tokenString
	glog.Infof("Auth passed, new user with token is: %v", user)

	resp := utils.Message(true, "Logged In")
	resp["user"] = user
	return resp
}

func GetUser(uid uint) *User {
	user := &User{}
	GetDb().Table("User").Where("id = ?", uid).First(user)

	if user.Email == "" { //User not found
		glog.Error("GetUser ID failed due to email is not in DB or DB connection issue")
		return nil
	}

	user.Password = ""
	return user
}
