package controllers

import (
	"net/http"
	"encoding/json"
	"github.com/atyu1/SSPro-Server/utils"
	"github.com/atyu1/SSPro-Server/models"
	"github.com/julienschmidt/httprouter"
)

/*
func DataPointHandler(res http.ResponseWriter, req *http.Request, param httprouter.Params) {
        // ToDo:
        // 1. Parse data
        // 2. Get data point
        // 3. Save to datapoint variable
        // 4. Return datapoint

        fmt.Fprintf(res, "test\nJOJO\n")
}

func DataPointGet(res http.ResponseWriter, req *http.Request, param httprouter.Params) {
        // Same todo

        roomname := param.ByName("room")
        fmt.Fprintf(res, "GO GET IT: %s!\n", roomname)
}
*/

func CreateDataPoint(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	tmpdatapoints := &datapoints.Datapoints{}
	
	err := json.NewDecoder(r.Body).Decode(tmpdatapoints)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error decoding the data in request body!"))
		return
	}

	resp := tmpdatapoints.Save()
	
	utils.Respond(w, resp)
}
