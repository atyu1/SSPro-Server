package main

import(
	"net/http"
	"fmt"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func main() {
	
	router := httprouter.New()
	router.GET("/test", Index)

	http.ListenAndServe(":8080", nil)
}
