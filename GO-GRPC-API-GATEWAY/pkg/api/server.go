package server

import (
	"log"

	"github.com/AnanyaDevGo/GO-GRPC-API-GATEWAY/pkg/api/handler"
	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(adminHandler *handler.AdminHandler) *ServerHTTP {
	router := gin.New()

	router.Use(gin.Logger())

	router.POST("/admin/login", adminHandler.LoginHandler)
	router.POST("/admin/signup", adminHandler.AdminSignUp)

	return &ServerHTTP{engine: router}
}

func (s *ServerHTTP) Start() {
	log.Printf("starting server on :3000")
	err := s.engine.Run(":6000")
	if err != nil {
		log.Printf("error while starting the server")
	}
}
