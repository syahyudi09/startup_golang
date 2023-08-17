package handler

import (
	"startup/config"
	"startup/manager"
	"startup/middleware"

	"github.com/gin-gonic/gin"
)

	type Server interface {
		Run()
	}

	type serverImpl struct {
		engine *gin.Engine
		usecase manager.UsecaseManager
	}

	func (s *serverImpl) Run() {
		NewUserHandler(s.engine, s.usecase.GetUserUsecase())
		s.engine.Run(":8080")

	}

	func NewServer() Server {
		config, err := config.NewConfig()
		if err != nil{
			panic(err)		
		}
		auth := middleware.NewJwtService()
		infra := manager.NewInfraManager(config)
		repo := manager.NewRepomanager(infra)
		usecase := manager.NewUsecasemanager(repo, auth)

		engine := gin.Default()

		return &serverImpl{
			engine: engine,
			usecase: usecase,
		}
	}