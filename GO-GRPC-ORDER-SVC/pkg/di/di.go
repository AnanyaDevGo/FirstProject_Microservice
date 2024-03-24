package di

import (
	server "github.com/AnanyaDevGo/pkg/api"
	"github.com/AnanyaDevGo/pkg/api/service"
	"github.com/AnanyaDevGo/pkg/client"
	"github.com/AnanyaDevGo/pkg/config"
	"github.com/AnanyaDevGo/pkg/db"
	"github.com/AnanyaDevGo/pkg/repository"
	"github.com/AnanyaDevGo/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*server.Server, error) {

	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	orderRepository := repository.NewOrderRepository(gormDB)
	cartClient := client.NewCartClient(&cfg)
	productClient := client.NewProductClient(&cfg)
	orderUseCase := usecase.NewOrderUseCase(orderRepository, cartClient, productClient)

	orderServiceServer := service.NewOrderServer(orderUseCase)
	grpcServer, err := server.NewGRPCServer(cfg, orderServiceServer)

	if err != nil {
		return &server.Server{}, err
	}

	return grpcServer, nil

}
