package manager

import (
	"startup/repo"
	"sync"
)

type RepoManager interface {
	GetUserRepo() repo.UserRepo
}

type repoManagerImpl struct {
	infra InfraManager
	userRepo repo.UserRepo
}

var onceLoadUserRepo sync.Once

func (rm *repoManagerImpl) GetUserRepo() repo.UserRepo{
	onceLoadUserRepo.Do(func() {
		rm.userRepo = repo.NewUserRepo(rm.infra.GetDB())
	})
	return rm.userRepo
}

func NewRepomanager(infra InfraManager) RepoManager {
	return &repoManagerImpl{
		infra: infra,
	}
}
