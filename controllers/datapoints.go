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


func GetDataPointAll(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	data, err := datapoints.GetDataPoints("")
	if err != nil {
		utils.Respond(w, utils.Message(false, err.Error()))
	}
	resp := utils.Message(true, "Data Collected")
	resp["data"] = data
	fmt.Printf("## %v\n", resp)
	utils.Respond(w, resp)
}


func CreateDataPoint(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	tmpdatapoints := &datapoints.Datapoints{}

	err := json.NewDecoder(r.Body).Decode(tmpdatapoints)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error decoding the data in request body!"))
		log.Print(fmt.Sprintf("%v", err))
		return
	}
	resp := tmpdatapoints.Save()
	utils.Respond(w, resp)
}
