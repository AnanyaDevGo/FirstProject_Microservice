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

	productClient := client.NewProductClient(cfg)
	productHandler := handler.NewProductHandler(productClient)

	userClient := client.NewUserClient(cfg)
	userHandler := handler.NewUserHandler(userClient)

	cartClient := client.NewCartClient(cfg)
	cartHandler := handler.NewCartHandler(cartClient)

	orderClient := client.NewOrderClient(cfg)
	orderHandler := handler.NewOrderHandler(orderClient)

	serverHTTP := server.NewServerHTTP(adminHandler, productHandler, userHandler, cartHandler, orderHandler)

	return serverHTTP, nil
}
