package main

import (
	"log"

	"github.com/Pratam-Kalligudda/product-service-go/config"
	"github.com/Pratam-Kalligudda/product-service-go/internal/api"
)

func main() {
	// will fill this later after completing config and server in api or rest
	appConfig, err := config.SetupConfig()
	if err != nil {
		log.Fatal(err)
	}
	api.SetupServer(appConfig)
}
