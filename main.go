package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

const HTDOCS = "htdocs"

func main() {
	if fi, err := os.Stat(HTDOCS); os.IsNotExist(err) {
		log.Fatalln("htdocs not exist")
	} else if !fi.IsDir() {
		log.Fatalln("htdocs is not dir")
	}

	addr := flag.String("addr", ":8080", "http server listen address, format: [ip]:port")
	flag.Parse()
	http.Handle("/", http.FileServer(http.Dir(HTDOCS)))
	log.Fatal(http.ListenAndServe(*addr, nil))
}
