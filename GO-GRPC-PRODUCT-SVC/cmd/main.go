package main

import (
	"log"

	"github.com/AnanyaDevGo/GO-GRPC-PRODUCT-SVC/pkg/config"
	"github.com/AnanyaDevGo/GO-GRPC-PRODUCT-SVC/pkg/di"
)

func main() {

	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}

	server, diErr := di.InitializeAPI(config)

	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start()
	}

}
