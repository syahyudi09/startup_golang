package manager

import (
	"startup/repo"
	"sync"
)

type RepoManager interface {
	GetUserRepo() repo.UserRepo
	GetCampaignRepo() repo.CampaignRepo
}

type repoManagerImpl struct {
	infra InfraManager
	userRepo repo.UserRepo
	campaignRepo repo.CampaignRepo
}

var onceLoadUserRepo sync.Once
var onceLoadCampignRepo sync.Once

func (rm *repoManagerImpl) GetUserRepo() repo.UserRepo{
	onceLoadUserRepo.Do(func() {
		rm.userRepo = repo.NewUserRepo(rm.infra.GetDB())
	})
	return rm.userRepo
}

func (rm *repoManagerImpl) GetCampaignRepo() repo.CampaignRepo{
	onceLoadCampignRepo.Do(func() {
		rm.campaignRepo = repo.NewCampaignRepo(rm.infra.GetDB())
	})
	return rm.campaignRepo
}

func NewRepomanager(infra InfraManager) RepoManager {
	return &repoManagerImpl{
		infra: infra,
	}
}
