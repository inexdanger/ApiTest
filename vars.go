package main

var (
	SelectPreguntasExamen = "select Pregunta,Respuesta1,Respuesta2,Respuesta3,Respuesta4,Correcta from preguntas where Idexamen =?;"
	InsertExamen          = "insert into examenes values (?,?);"
	InserPreguntas        = "insert into preguntas(IdExamen,Pregunta,Respuesta1,Respuesta2,Respuesta3,Respuesta4,Correcta) values (?,?,?,?,?,?,?)"
)

type DataExamen struct {
	Preguntas []Pregunta
	Mensaje   string
}
type Pregunta struct {
	ID         int
	Pregunta   string
	Respuesta1 string
	Respuesta2 string
	Respuesta3 string
	Respuesta4 string
	Correcta   string
}
type Data struct {
	ID        string
	Nombre    string
	Preguntas string
	Inputs    []int
}
