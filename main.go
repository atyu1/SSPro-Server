package main

import(
	"net/http"
	"log"
	"github.com/atyu1/SSPro-Server/datapoints"
	"github.com/julienschmidt/httprouter"
)


func main() {

	var dataPointHandler = datapoints.DataPointHandler
	var dataPointGet = datapoints.DataPointGet

	router := httprouter.New()
	router.GET("/datapoints/:location", dataPointGet)
	router.GET("/datapoints/:location/:room", dataPointGet)
	router.GET("/datapoints/:location/:room/:name", dataPointGet)

	router.POST("/datapoints", 
	

	log.Fatal(http.ListenAndServe(":8080", router))
}
