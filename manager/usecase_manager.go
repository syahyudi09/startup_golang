package manager

import (
	"startup/middleware"
	"startup/usecase"
	"sync"
)

type UsecaseManager interface {
	GetUserUsecase() usecase.Userusecase
	GetCampaignUsecase() usecase.CampaignUsecase
}

type usecaseManager struct {
	repoManager  RepoManager
	userUsecase usecase.Userusecase
	campaignUsecaase usecase.CampaignUsecase
	auth middleware.Auth
}

var onceLoadUserUsecase sync.Once
var onceLoadCampaignUsecase sync.Once

func (um *usecaseManager) GetUserUsecase() usecase.Userusecase{ 
	onceLoadUserUsecase.Do(func ()  {
		um.userUsecase = usecase.NewUserUsecase(
			um.repoManager.GetUserRepo(),
			um.auth,
		)
	})
	return um.userUsecase 
}

func (um *usecaseManager) GetCampaignUsecase() usecase.CampaignUsecase{
	onceLoadCampaignUsecase.Do(func() {
		um.campaignUsecaase = usecase.NewCampaignUsecase(
			um.repoManager.GetCampaignRepo(),
			um.auth,
		)
	})
	return um.campaignUsecaase
}

func NewUsecasemanager(repo RepoManager, auth middleware.Auth) UsecaseManager {
	return &usecaseManager{
		repoManager: repo,
		auth: auth,
		
	}
}