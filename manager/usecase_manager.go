package manager

import (
	"startup/usecase"
	"sync"
)

type UsecaseManager interface {
	GetUserUsecase() usecase.Userusecase
}

type usecaseManager struct {
	repoManager  RepoManager
	userUsecase usecase.Userusecase
}

var onceLoadUserUsecase sync.Once

func (um *usecaseManager) GetUserUsecase() usecase.Userusecase{ 
	onceLoadUserUsecase.Do(func ()  {
		um.userUsecase = usecase.NewUserUsecase(um.repoManager.GetUserRepo())
	})
	return um.userUsecase
}

func NewUsecasemanager(repo RepoManager) UsecaseManager {
	return &usecaseManager{
		repoManager: repo,
	}
}