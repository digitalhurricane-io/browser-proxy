package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", proxyHandler)

	log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), nil))
}