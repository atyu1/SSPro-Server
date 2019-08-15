package main

import (
	"github.com/atyu1/SSPro-Server/controllers"
	"github.com/atyu1/SSPro-Server/models"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {

	//ToDo: Change to ENV or so
	CONFIG_FILE := "./config.yaml"

	config := datapoints.DbAccess{}
	config.Init(CONFIG_FILE)

	//sqlite3 used
	datapoints.InitDb(config)

	router := httprouter.New()
	router.GET("/datapoints/all", controllers.GetDataPointAll)
	//router.GET("/datapoints/:location", dataPointGet)
	//router.GET("/datapoints/:location/:room", dataPointGet)
	//router.GET("/datapoints/:location/:room/:name", dataPointGet)

	router.POST("/datapoints", controllers.CreateDataPoint)

	log.Fatal(http.ListenAndServe(":8080", router))
	datapoints.GetDb().Close()
}
