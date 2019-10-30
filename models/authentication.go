package authentication

import (
	"strings"
	"golang.org/x/crypto/bcrypt"
	"github.com/atyu/SSPro-Server/utils"
	"github.com/jinzhu/gorm"
	"github.com/dgrijalva/jwt-go"
)

// JWT User struct
type Token struct {
	UserId uint
	jwt.StandardClaims
}


// User is a struct to define User
type User struct {
	gorm.Model
	Email string `json:"email"`
	Password string `json:"password"`
	Token string `json:"token";sql:"-"`
}

func Login(email, password, tokenPassword string) (map[string]interface{}) {
	user := &User{}

	err := GetDB().Table("users").Where("email = ?",email).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound{
			return utils.Message(false, "Email address not found")
		}
		return utils.Message(fase, "Database connection error")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.password), []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return utils.Message(false, "Invalid login credentials")
	}

	user.Password = ""  //Remove password from memory

	token := &Token{UserrId: user.ID}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"),token)
	tokenString, _ := jwtToken.SignedString([]byte(tokenPassword))
	user.Token = tokenString

	resp := utils.Message(true, "Logged In")
	resp["account"] = account
	return resp
}

func GetUser(uid uint) *User {
	user := &User{}
	GetDB.Table("User").Where("id = ?", uid).First(user)

	if user.Email == "" { //User not found
		return nil
	}
	
	user.Password = ""
	return user
}
