	package handler

	import (
		"startup/config"
		"startup/manager"

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

		infra := manager.NewInfraManager(config)
		repo := manager.NewRepomanager(infra)
		usecase := manager.NewUsecasemanager(repo)

		engine := gin.Default()

		return &serverImpl{
			engine: engine,
			usecase: usecase,
		}
	}