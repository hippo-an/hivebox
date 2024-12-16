package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	appVersion = "v0.0.1"
	port       = 8888
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(appVersion))
	})

	log.Printf("application server running on port %d\n", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
