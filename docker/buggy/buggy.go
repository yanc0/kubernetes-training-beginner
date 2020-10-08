package main

import (
	"io"
	"log"
	"net/http"
)

func fail(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	io.WriteString(w, "500, FATAL - End of the world !")
}

func main() {
	http.HandleFunc("/", fail)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
