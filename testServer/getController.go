package main

import (
	"encoding/json"
	"net/http"
)

var client *http.Client

type CatImg struct {
	Url string `json:"url"`
}

func getCatImg(wr http.ResponseWriter, r *http.Request){
	url := "https://api.thecatapi.com/v1/images/search?0709efa2-9cd7-42ea-a70f-a04237af5797"

	var catImg CatImg

	err := getJson(url, &catImg)

	if err != nil {
		panic(err.Error)
	}
		wr.Write([]byte(catImg.Url))
	

}

func getJson(url string, target interface{}) error{
	resp, err := client.Get(url)

	if err != nil {
		panic(err.Error)
	}
	
	defer resp.Body.Close()
	
	return json.NewDecoder(resp.Body).Decode(target)
}