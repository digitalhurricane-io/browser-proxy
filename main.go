package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", proxyHandler)

	log.Fatal(http.ListenAndServe(":9100", nil))
}