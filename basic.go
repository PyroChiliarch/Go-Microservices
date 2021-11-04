package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Receiving Data!")
		d, err := io.ReadAll(r.Body)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write(([]byte("oops")))
		}

		fmt.Fprintf(rw, "Hello %s", d)
	})

	http.HandleFunc("/helicopter", func(http.ResponseWriter, *http.Request) {
		log.Println("HelicopterGoBRrRrRrRRrRRrrRRrRRRrRrrR!!!")
	})

	http.HandleFunc("/bad", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Failure")
		http.Error(rw, "Blap, you failed", http.StatusBadRequest)
	})

	http.ListenAndServe(":9090", nil)
}
