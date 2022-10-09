package main

import (
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
)

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
    fmt.Fprintf(w,"Home handler")
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Products handler")
}

func ArticlesHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Category handler")
}

