package main

import (
	"handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)

	/*
		http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {

		})

		http.HandleFunc("/helicopter", func(http.ResponseWriter, *http.Request) {
			log.Println("HelicopterGoBRrRrRrRRrRRrrRRrRRRrRrrR!!!")
		})

		http.HandleFunc("/bad", func(rw http.ResponseWriter, r *http.Request) {
			log.Println("Failure")
			http.Error(rw, "Blap, you failed", http.StatusBadRequest)
		})
	*/
	http.ListenAndServe(":9090", nil)

}
