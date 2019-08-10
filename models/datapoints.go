package datapoints

import (
	"net/http"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/atyu1/SSPro-Server/utils"
)


type Datapoint struct {
	Location string  `json:"location"`
	Room     string  `json:"room"`
	Name     string  `json:"name"`
	Sensor   string  `json:"sensor"`
	Value    float64 `json:"value"`
}

func (d *Datapoint) Validate() (map[string]interface{}, bool) {
	
	if d.Location == "" {
		return u.Message(false, "Location is empty"), false
	}

	if d.Room == "" {
		return u.Message(false, "Room is empty"), false
	}

	if d.Name == "" {
		return u.Message(false, Name is empty"), false
	}

	if d.Sensor == "" {
		return u.Message(false, Sensor is empty"), false
	}

	return u.Message(true, "success"), true
}


func (d *Datapoint) Save (

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
