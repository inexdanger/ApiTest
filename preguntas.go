package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func grabacionPreguntas(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		Preguntas(w, r)
	case "POST":
		PreguntasPost(w, r)
	}

}

func Preguntas(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	db, err := connectDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	funcMap := template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
	}
	data := Data{
		ID:        id,
		Preguntas: r.FormValue("preguntas"),
		Inputs:    make([]int, 0),
	}

	numPreguntas, err := strconv.Atoi(data.Preguntas)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < numPreguntas; i++ {
		data.Inputs = append(data.Inputs, i)
	}

	tmpl := template.Must(template.New("examen2.html").Funcs(funcMap).ParseFiles("examen2.html"))
	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func PreguntasPost(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	preguntas, err := strconv.Atoi(r.FormValue("preguntas"))
	if err != nil {
		fmt.Println(err)
		return
	}
	db, err := connectDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	o := 0

	for i := 0; i < preguntas; i++ {
		_, err = db.Query(InserPreguntas, id, r.FormValue("pregunta"+strconv.Itoa(i)), r.FormValue("respuesta"+strconv.Itoa(i)+strconv.Itoa(o+1)), r.FormValue("respuesta"+strconv.Itoa(i)+strconv.Itoa(o+2)), r.FormValue("respuesta"+strconv.Itoa(i)+strconv.Itoa(o+3)), r.FormValue("respuesta"+strconv.Itoa(i)+strconv.Itoa(o+4)), r.FormValue("true"+strconv.Itoa(i)))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	w.Write([]byte("<script>alert(" + id + ")</script>"))
}
