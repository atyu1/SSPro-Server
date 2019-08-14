package datapoints

import (
	"fmt"
	"github.com/atyu1/SSPro-Server/utils"
	//"github.com/jinzhu/gorm"
)

type Datapoint struct {
	//gorm.Model
	Location string  `json:"location"`
	Room     string  `json:"room"`
	Name     string  `json:"name"`
	Sensor   string  `json:"sensor"`
	Value    float64 `json:"value"`
}

type Datapoints struct {
	Data []Datapoint `json:"data"`
}

/*
Validate function checks the required parameters in http request body and checks if data is not missing

returns message and true if all the data is correct and not missing
*/
func (d *Datapoint) Validate() (map[string]interface{}, bool) {

	if d.Location == "" {
		return utils.Message(false, "Location is empty"), false
	}

	if d.Room == "" {
		return utils.Message(false, "Room is empty"), false
	}

	if d.Name == "" {
		return utils.Message(false, "Name is empty"), false
	}

	if d.Sensor == "" {
		return utils.Message(false, "Sensor is empty"), false
	}

	return utils.Message(true, "success"), true
}

/*
Save datapoint to the database

returns message about success
*/
func (ds *Datapoints) Save() map[string]interface{} {

	for _, d := range ds.Data {
		if _, ok := d.Validate(); ok {
			fmt.Printf("\n-- %v\n", d)
			GetDb().Create(d)
		}
	}

	resp := utils.Message(true, "DataPoints created")
	return resp
}
