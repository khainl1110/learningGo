package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"html/template"
)

// declaring struct
type Action struct {
	// defining struct fields
	Name   string
	Action string
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/products", ProductsHandler)
	r.HandleFunc("/articles", ArticlesHandler)
	http.Handle("/", r)

	http.ListenAndServe(":80", r)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)

	std1 := Action{"Vani", "Home handler"}

	t, _ := template.ParseFiles("templates/template.html")

	// web output to print merged data
	t.Execute(w, std1)

}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)

	std2 := Action{"Vani", "Products handler"}

	t, _ := template.ParseFiles("templates/template.html")

	t.Execute(w, std2)
}

func ArticlesHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)

	std1 := Action{"Vani", "Articles handler"}

	t, _ := template.ParseFiles("templates/template.html")

	t.Execute(w, std1)
}
