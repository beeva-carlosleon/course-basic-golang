package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"strconv"
)

const responseBody = "hello, BEEVA!\n"

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, responseBody)
}

func launchServer(port *int) {
	log.Printf("Listening on port: '%d'", *port)
	http.HandleFunc("/hello", HelloServer)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*port), nil))
}
func main() {
	port := flag.Int("port", 12345, "Port to listen")
	flag.Parse()
	launchServer(port)
}
