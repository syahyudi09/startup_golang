package manager

import (
	"startup/middleware"
	"startup/usecase"
	"sync"
)

type UsecaseManager interface {
	GetUserUsecase() usecase.Userusecase
}

type usecaseManager struct {
	repoManager  RepoManager
	userUsecase usecase.Userusecase
	auth middleware.Auth
}

var onceLoadUserUsecase sync.Once

func (um *usecaseManager) GetUserUsecase() usecase.Userusecase{ 
	onceLoadUserUsecase.Do(func ()  {
		um.userUsecase = usecase.NewUserUsecase(
			um.repoManager.GetUserRepo(),
			um.auth,
		)
	})
	return um.userUsecase 
}

func NewUsecasemanager(repo RepoManager, auth middleware.Auth) UsecaseManager {
	return &usecaseManager{
		repoManager: repo,
		auth: auth,
		
	}
}