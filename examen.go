package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
)

func Examen(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ExamenGet(w, r)
	case "POST":
		ExamenPost(w, r)
	}
}

func ExamenGet(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("examen.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ExamenPost(w http.ResponseWriter, r *http.Request) {
	nombre := r.FormValue("nombre")
	preguntas := r.FormValue("preguntas")

	id := rand.Intn(100000)

	http.Redirect(w, r, fmt.Sprintf("/preguntas/%d?id=%d&nombre=%s&preguntas=%s", id, id, nombre, preguntas), http.StatusFound)
}

func GetExamen(w http.ResponseWriter, r *http.Request) {
	var examen DataExamen
	vars := mux.Vars(r)
	id := vars["id"]
	examen.Mensaje = id
	db, err := connectDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	rows, err := db.Query(SelectPreguntasExamen, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var pregunta Pregunta
		err = rows.Scan(
			&pregunta.Pregunta,
			&pregunta.Respuesta1,
			&pregunta.Respuesta2,
			&pregunta.Respuesta3,
			&pregunta.Respuesta4,
			&pregunta.Correcta,
		)
		if err != nil {
			fmt.Println(err)
			return
		}
		examen.Preguntas = append(examen.Preguntas, pregunta)
	}

	tmpl := template.Must(template.ParseFiles("getExamen.html"))

	err = tmpl.Execute(w, examen)
	if err != nil {
		fmt.Println(err)
		return
	}
}
