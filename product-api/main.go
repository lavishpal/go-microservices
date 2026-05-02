package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello world")
		d, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Bad request", http.StatusBadRequest)
			//rw.WriteHeader(http.StatusBadRequest)
			//rw.Write([]byte("Bad request"))
			return

		}

		fmt.Fprintf(rw, "Hello %s", d)

		//log.Printf("Data %s\n", d)

	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("GoodBye world")
	})

	http.ListenAndServe(":9090", nil)

}
