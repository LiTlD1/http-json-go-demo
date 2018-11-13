package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

//Return string from url query function
func makeURLRequest(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return string(body)
}
