package main

import "net/http"


func main(){
	// http.HandleFunc("/", helloWorld)
	http.HandleFunc("/", getCatImg)
	http.ListenAndServe(":8000", nil)
	
}