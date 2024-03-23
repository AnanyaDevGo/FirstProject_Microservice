package di

import (
	server "github.com/AnanyaDevGo/GO-GRPC-PRODUCT-SVC/pkg/api"
	"github.com/AnanyaDevGo/GO-GRPC-PRODUCT-SVC/pkg/api/service"
	"github.com/AnanyaDevGo/GO-GRPC-PRODUCT-SVC/pkg/config"
	"github.com/AnanyaDevGo/GO-GRPC-PRODUCT-SVC/pkg/db"
	"github.com/AnanyaDevGo/GO-GRPC-PRODUCT-SVC/pkg/repository"
	"github.com/AnanyaDevGo/GO-GRPC-PRODUCT-SVC/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*server.Server, error) {

	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	adminRepository := repository.NewProductRepository(gormDB)
	adminUseCase := usecase.NewProductUseCase(adminRepository)

	adminServiceServer := service.NewProductServer(adminUseCase)
	grpcServer, err := server.NewGRPCServer(cfg, adminServiceServer)

	if err != nil {
		return &server.Server{}, err
	}

	return grpcServer, nil

}
