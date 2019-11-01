package authentication

import (
	"net/http"
	"strings"
	"fmt"
	"github.com/golang/glog"
	"github.com/atyu1/SSPro-Server/utils"
	"github.com/julienschmidt/httprouter"
 	"github.com/atyu1/SSPro-Server/models"
	jwt "github.com/dgrijalva/jwt-go"
)

var TOKENPASS string = "test" 

// JWTAuth is a custom handler which first verify if user has Token and it is authorized for API calls
// We expect in HTTP header to have: Authorization: Bearer <token>
func JWTAuth(h httprouter.Handle) (httprouter.Handle) {
	return func(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
		tokenHeader := r.Header.Get("Authorization")
		glog.Infof("TokenHeader: %s",tokenHeader)
		// If Token is missing
		if tokenHeader == "" {
			statusForbiden("Missing Authorization Header", w)
			return
		}

		// If Token has invalid format
		// Expected to see in JWT tokens "Bearer <token>"
		splittedToken := strings.Split(tokenHeader, " ")
		if len(splittedToken) != 2 {
			statusForbiden("Authorization header is not in expected format", w)
			return
		}

		tokenTmp := splittedToken[1] // Remove Bearer
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenTmp, tk, func(token *jwt.Token) (interface{}, error) {
			return[]byte(TOKENPASS), nil
		})

		if err != nil {
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors & jwt.ValidationErrorMalformed != 0 {
					statusForbiden("This is not token, token is totally malformed",w)
				} else if ve.Errors & (jwt.ValidationErrorExpired | jwt.ValidationErrorNotValidYet) != 0 {
					statusForbiden("Token is either expired or not valid yet!", w)
				} else {
					statusForbiden("Issues with the provided token", w)
				}
			} else {
				statusForbiden("Issues with the provided token", w)
			}
			return 
		}

		if !token.Valid {
			statusForbiden("Token validation failed", w)
			return
		}

		//Token is valid finally!
		h(w, r, param)	
	}
}

// statusForbiden is used in any authentication issues related problems
// We call with text and http response writer to send back a specific issue with 403 code
func statusForbiden(text string, w http.ResponseWriter) {
	response := utils.Message(false, text)
	w.WriteHeader(http.StatusForbidden)
	w.Header().Add("Content-Type", "application/json")
	utils.Respond(w, response)
}
