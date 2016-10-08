package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Printf("starting the web server")

	http.HandleFunc("/", hello)
	http.ListenAndServe(":", nil)
}

func hello(rw http.ResponseWriter, r *http.Request) {
	log.Printf("%s from %s", r.Method, r.RemoteAddr)
	fmt.Fprint(rw, "Hello from the other side")
}
