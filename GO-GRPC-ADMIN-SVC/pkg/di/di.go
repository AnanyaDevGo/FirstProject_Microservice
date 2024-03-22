package di

import (
	server "github.com/AnanyaDevGo/GO-GRPC-ADMIN-SVC/pkg/api"
	"github.com/AnanyaDevGo/GO-GRPC-ADMIN-SVC/pkg/api/service"
	"github.com/AnanyaDevGo/GO-GRPC-ADMIN-SVC/pkg/config"
	"github.com/AnanyaDevGo/GO-GRPC-ADMIN-SVC/pkg/db"
	"github.com/AnanyaDevGo/GO-GRPC-ADMIN-SVC/pkg/repository"
	"github.com/AnanyaDevGo/GO-GRPC-ADMIN-SVC/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*server.Server, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	adminRepository := repository.NewAdminRepository(gormDB)
	adminUseCase := usecase.NewAdminUseCase(adminRepository)
	adminServiceServer := service.NewAdminServer(adminUseCase)
	grpcServer, err := server.NewGRPCServer(cfg, adminServiceServer)

	if err != nil {
		return &server.Server{}, err
	}
	return grpcServer, nil
}
