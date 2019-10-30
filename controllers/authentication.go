package authentication

import (
	"net/http"
	"encoding/json"
	"github.com/atyu1/SSPro-Server/models"
        "github.com/atyu1/SSPro-Server/utils"
        "github.com/julienschmidt/httprouter"
)

var tokenPassword string = "ThisIsTokenPasswordTemporary"

//Login is a function to login through HTTP parameters
func Login(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	user := &authentication.User{} //From models
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		return utils.Message(false, "Invalid login request")
	}
	resp := authentication.Login(user.Email, user.Password, tokenPassword)
	utils.Respond(w, resp)
}
