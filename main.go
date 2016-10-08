package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Printf("starting the web server")

	http.HandleFunc("/", hello)
	http.ListenAndServe(":80", nil)
}

func hello(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Hello "+r.RemoteAddr)
}
