package main

import(
	"net/http"
	"fmt"
	"datapoints/datapoints"
)


func main() {

	var d datapoint
	
	router := http.NewServeMux()
	router.Handle("/test", d)

	http.ListenAndServe(":8080", router)
}
