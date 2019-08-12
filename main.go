package main

import(
	"net/http"
	"log"
	"github.com/atyu1/SSPro-Server/models" 
	"github.com/atyu1/SSPro-Server/controllers" 
	"github.com/julienschmidt/httprouter"
)


func main() {

	//ToDo: Change to ENV or so
	CONFIG_FILE := "github.com/atyu1/SSPro-Server/config.yaml"

	config := datapoints.DbAccess{}
	config.Init(CONFIG_FILE)

	//sqlite3 used
	datapoints.InitDb(config)

	router := httprouter.New()
	//router.GET("/datapoints/:location", dataPointGet)
	//router.GET("/datapoints/:location/:room", dataPointGet)
	//router.GET("/datapoints/:location/:room/:name", dataPointGet)

	router.POST("/datapoints", controllers.CreateDataPoint)
	

	log.Fatal(http.ListenAndServe(":8080", router))
	datapoints.GetDb().Close()
}
