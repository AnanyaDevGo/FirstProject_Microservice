package di

import (
	server "github.com/AnanyaDevGo/pkg/api"
	"github.com/AnanyaDevGo/pkg/api/service"
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

	adminRepository := repository.NewUserRepository(gormDB)
	adminUseCase := usecase.NewUserUseCase(adminRepository)

	userServiceServer := service.NewUserServer(adminUseCase)
	grpcServer, err := server.NewGRPCServer(cfg, userServiceServer)

	if err != nil {
		return &server.Server{}, err
	}

	return grpcServer, nil

}
