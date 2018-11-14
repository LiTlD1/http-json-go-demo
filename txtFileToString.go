package main

import (
	"io/ioutil"
	"log"
)

func txtFileToString(fileString string) string {
	b, err := ioutil.ReadFile(fileString)
	if err != nil {
		log.Fatalln(err)
	}

	return string(b)
}
