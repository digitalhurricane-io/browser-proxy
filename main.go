package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", proxyHandler)

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request){
		setCors(&w)
		if r.Method == "OPTIONS" {
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), nil))
}