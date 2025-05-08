package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hippo-an/hivebox/config"
	"github.com/hippo-an/hivebox/handlers"
)

func main() {
	config.InitConfig()
	config.InitSecret()

	http.HandleFunc("GET /version", handlers.GetVersion)
	http.HandleFunc("GET /forecast/temperature", handlers.GetForecastTemperatureHandler)

	port := config.AppConfig.Application.Port
	log.Printf("hivebox server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
