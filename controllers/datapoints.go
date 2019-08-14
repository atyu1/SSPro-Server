package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"github.com/atyu1/SSPro-Server/models"
	"github.com/atyu1/SSPro-Server/utils"
	"github.com/julienschmidt/httprouter"
	"net/http"
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
		log.Print(fmt.Sprintf("%v", err))
		return
	}
	resp := tmpdatapoints.Save()
	fmt.Printf("We are here")
	utils.Respond(w, resp)
}
