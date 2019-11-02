package controllers

import (
	"encoding/json"
	"github.com/atyu1/SSPro-Server/models"
	"github.com/atyu1/SSPro-Server/utils"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/golang/glog"
)

//Login is a function to login through HTTP parameters
func Login(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	user := &models.User{} //From models
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		glog.Errorf("Decoding json body issues: %v", err)
		utils.Respond(w, utils.Message(false, "Invalid request"))
	} else {
		glog.Infof("Provided credentials: %v", user)
		resp := models.Login(user.Email, user.Password)
		utils.Respond(w, resp)
	}
}
