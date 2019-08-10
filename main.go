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
	router.GET("/test", dataPointHandler)
	router.GET("/datapoints/:room", dataPointGet)

	log.Fatal(http.ListenAndServe(":8080", router))
}
