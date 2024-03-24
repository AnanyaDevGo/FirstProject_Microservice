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

	cartRepository := repository.NewCartRepository(gormDB)
	productClient := client.NewProductClient(&cfg)
	adminUseCase := usecase.NewCartUseCase(cartRepository, productClient)

	adminServiceServer := service.NewCartServer(adminUseCase)
	grpcServer, err := server.NewGRPCServer(cfg, adminServiceServer)

	if err != nil {
		return &server.Server{}, err
	}

	return grpcServer, nil

}
