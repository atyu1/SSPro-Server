package datapoint

type datapoint struct {
	Location string  `json:"location"`
	Room     string  `json:"room"`
	Name     string  `json:"name"`
	Sensor   string  `json:"sensor"`
	Value    float64 `json:"value"`
}


