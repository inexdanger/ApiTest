package main

import (
	"database/sql"
	"fmt"
)

func connectDB() (db *sql.DB, e error) {
	usuario := "root"
	pass := "temporal"
	host := "tcp(127.0.0.1:3306)"
	nombreBaseDeDatos := "testApi"
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", usuario, pass, host, nombreBaseDeDatos))
	if err != nil {
		return nil, err
	}
	return db, nil
}
