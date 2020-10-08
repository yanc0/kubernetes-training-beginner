package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	msg := os.Getenv("MSG")
	if msg == "" {
		msg = "Hello world !"
	}
	io.WriteString(w, fmt.Sprintf("%s - %s\n", hostname, msg))
}

func main() {
	http.HandleFunc("/", hello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
