package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hippo-an/hivebox/handlers"
)

const (
	port = 8888
)

func main() {
	http.HandleFunc("GET /version", handlers.GetVersion)
	http.HandleFunc("GET /forecast/temperature", handlers.GetForecastTemperatureHandler)

	log.Printf("hivebox server running on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
