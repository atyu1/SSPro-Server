package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/atyu1/SSPro-Server/models"
	"github.com/atyu1/SSPro-Server/utils"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strings"
)

func GetDataPointAll(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	tmp := param.ByName("location")
	params := strings.Split(tmp, "/")

	data, err := models.GetDataPoints([]int{}, params)
	if err != nil {
		utils.Respond(w, utils.Message(false, err.Error()))
		return
	}
	resp := utils.Message(true, "Data Collected")
	resp["data"] = data
	utils.Respond(w, resp)
}

func CreateDataPoint(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	tmpdatapoints := &models.Datapoints{}

	err := json.NewDecoder(r.Body).Decode(tmpdatapoints)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error decoding the data in request body!"))
		log.Print(fmt.Sprintf("%v", err))
		return
	}
	resp := tmpdatapoints.Save()
	utils.Respond(w, resp)
}
