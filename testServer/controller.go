package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Person struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
	Height int `json:"height"`
}

func dbConnect() (db *sql.DB){
	
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "password"
	dbName := "test_go"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func helloWorld(wr http.ResponseWriter, r *http.Request){
	db := dbConnect()
	selDB, err := db.Query("SELECT * FROM person")

	defer db.Close()
	defer selDB.Close()

	if err != nil {
		panic(err.Error())
	}
	
	person := Person{}
	res := []Person{}

	for selDB.Next() {
		var id, age, height int
		var name string
		
		err = selDB.Scan(&id, &name, &age, &height)

		if err != nil {
			panic(err.Error())
		}

		person.Id = id
		person.Name = name
		person.Age = age
		person.Height = height


		res = append(res, person)

	}

	jsonData, err := json.Marshal(res)
	rawIn := json.RawMessage(jsonData)

	if err != nil {
		http.Error(wr, err.Error(), http.StatusBadRequest)
		return
	}

	if err != nil {
		panic(err.Error())
	}

	newString := json.NewEncoder(wr).Encode(&rawIn)

	if err != nil {
		http.Error(wr, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(newString)
	

	defer db.Close()
}

