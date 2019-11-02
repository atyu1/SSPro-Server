package models

import (
	"github.com/atyu1/SSPro-Server/utils"
)

type Datapoint struct {
	Timestamp int     `json:"timestamp"`
	Location  string  `json:"location"`
	Room      string  `json:"room"`
	Name      string  `json:"name"`
	Sensor    string  `json:"sensor"`
	Value     float64 `json:"value"`
}

type Datapoints struct {
	Data []Datapoint `json:"data"`
}

/*
Validate function checks the required parameters in http request body and checks if data is not missing

returns message and true if all the data is correct and not missing
*/
func (d *Datapoint) Validate() (map[string]interface{}, bool) {
	if d.Timestamp <= 0 {
		return utils.Message(false, "Location is empty"), false
	}
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
			// Todo Add log for Validation results
			GetDb().Create(d)
		}
	}

	resp := utils.Message(true, "DataPoints created")
	return resp
}

/*
GetDataPoints function query the database and getting datapoints based on varariables

returns list of datapoints which fetch the query
*/
func GetDataPoints(query []int, params []string) (Datapoints, error) {
	var err error
	var (
		location string
		room     string
		name     string
	)

	for index, value := range params {
		switch index {
		case 1:
			location = value
		case 2:
			room = value
		case 3:
			name = value
		}
	}

	datapoints := []Datapoint{}

	err = getAllDatapoints(&datapoints, location, room, name)

	return Datapoints{Data: datapoints}, err
}

/// ----------------------- Helper functions --------------
/*getAllDatapoints is a helper fucntion for GetDataPoints used to get All (not per timestamp) datapoints for specific location, room, name

returns error if datapoints are not collected,
*/
func getAllDatapoints(dbpoints *[]Datapoint, location string, room string, name string) error {
	params := &Datapoint{Location: location, Room: room, Name: name}

	err := GetDb().Table("Datapoints").Where(params).Find(&dbpoints).Error
	return err
}
