package di

import (
	server "github.com/AnanyaDevGo/GO-GRPC-API-GATEWAY/pkg/api"
	"github.com/AnanyaDevGo/GO-GRPC-API-GATEWAY/pkg/api/handler"
	"github.com/AnanyaDevGo/GO-GRPC-API-GATEWAY/pkg/client"
	"github.com/AnanyaDevGo/GO-GRPC-API-GATEWAY/pkg/config"
)

func InitializeAPI(cfg config.Config) (*server.ServerHTTP, error) {

	adminClient := client.NewAdminClient(cfg)
	adminHandler := handler.NewAdminHandler(adminClient)

	serverHTTP := server.NewServerHTTP(adminHandler)

	return serverHTTP, nil
}
