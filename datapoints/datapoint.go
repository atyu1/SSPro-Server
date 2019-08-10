package datapoint

type datapoint struct {
	Location string  `json:"location"`
	Room     string  `json:"room"`
	Name     string  `json:"name"`
	Sensor   string  `json:"sensor"`
	Value    float64 `json:"value"`
}

func (d datapoint) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// ToDo:
	// 1. Parse data
	// 2. Get data point
	// 3. Save to datapoint variable
	// 4. Return datapoint

	io.WriteString(res, "test")
}
