package datapoints

import (
	"net/http"
	"fmt"
	"github.com/julienschmidt/httprouter"
)

type Datapoint struct {
	Location string  `json:"location"`
	Room     string  `json:"room"`
	Name     string  `json:"name"`
	Sensor   string  `json:"sensor"`
	Value    float64 `json:"value"`
}

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
