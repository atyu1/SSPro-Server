package main

import(
	"html/template"
	"net/http"
	"fmt"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("./templates/index.gohtml"))
	fmt.Println("Init done")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func main() {
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)
}
