package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/examen", Examen).Methods("GET", "POST")
	r.HandleFunc("/examen/{id}", GetExamen).Methods("GET", "POST")
	r.HandleFunc("/preguntas/{id}", grabacionPreguntas).Methods("GET", "POST")

	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)
}
