package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strings"
)

const HTDOCS = "htdocs"

func noDirListing(h http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func main() {
	if fi, err := os.Stat(HTDOCS); os.IsNotExist(err) {
		log.Fatalln("htdocs not exist")
	} else if !fi.IsDir() {
		log.Fatalln("htdocs is not dir")
	}

	addr := flag.String("addr", ":8080", "http server listen address, format: [ip]:port")
	enableDir := flag.Bool("enableDir", false, "enable list dir")
	flag.Parse()
	if *enableDir {
		http.Handle("/", http.FileServer(http.Dir(HTDOCS)))
	} else {
		http.Handle("/", noDirListing(http.FileServer(http.Dir(HTDOCS))))
	}
	log.Fatal(http.ListenAndServe(*addr, nil))
}
