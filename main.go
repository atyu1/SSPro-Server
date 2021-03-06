package main

import (
	"flag"
	"github.com/atyu1/SSPro-Server/auth"
	"github.com/atyu1/SSPro-Server/controllers"
	"github.com/atyu1/SSPro-Server/models"
	"github.com/golang/glog"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {

	flag.Parse()
	glog.Info("Program started")
	//ToDo: Change to ENV or so
	CONFIG_FILE := "./config.yaml"

	config := models.DbAccess{}
	config.Init(CONFIG_FILE)

	//sqlite3 used
	models.InitDb(config)

	router := httprouter.New()
	router.POST("/login", controllers.Login)
	router.GET("/datapoints/all/*location", authentication.JWTAuth(controllers.GetDataPointAll))
	router.POST("/datapoints", authentication.JWTAuth(controllers.CreateDataPoint))

	glog.Fatal(http.ListenAndServe(":8080", router))
	models.GetDb().Close()
}
