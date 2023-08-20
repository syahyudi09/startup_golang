package handler

import (
	"startup/config"
	"startup/manager"
	"startup/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type Server interface {
	Run()
}

type serverImpl struct {
	engine        *gin.Engine
	usecase       manager.UsecaseManager
	auth          middleware.Auth
}

func (s *serverImpl) Run() {
	auth := middleware.NewJwtService()
	middleware := NewMiddleware(auth)

	authenticated := s.engine.Group("/")
	authenticated.Use(middleware.AuthMiddleware())

	NewUserHandler(s.engine, s.usecase.GetUserUsecase())
	NewCampaignHandler(s.engine, s.usecase.GetCampaignUsecase())

	s.engine.Run(":8080")
}


func NewServer() Server {
	config, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	auth := middleware.NewJwtService()
	infra := manager.NewInfraManager(config)
	repo := manager.NewRepomanager(infra)
	usecase := manager.NewUsecasemanager(repo, auth)

	engine := gin.Default()

	store := cookie.NewStore([]byte(middleware.SECRET_KEY))
	engine.Use(sessions.Sessions("mysession", store))

	return &serverImpl{
		engine:        engine,
		usecase:       usecase,
		auth:          auth, // Inisialisasi authMiddleware yang Anda gunakan
	}
}
