package usecase

import (
	"startup/middleware"
	"startup/model"
	"startup/repo"
)

type CampaignUsecase interface {
	FindByID(int) (*model.CampaignModel, error)
	FindAll() ([]*model.CampaignModel, error)
}

type campaignUsecaseImpl struct {
	campaignRepo repo.CampaignRepo
	auth middleware.Auth
}

func (cu *campaignUsecaseImpl) FindByID(userID int) (*model.CampaignModel, error){
	campaign, err := cu.campaignRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}
	return campaign, nil
}

func (cu *campaignUsecaseImpl) FindAll() ([]*model.CampaignModel, error){
	return cu.campaignRepo.FindAll()
}


func NewCampaignUsecase(campaignRepo repo.CampaignRepo, auth middleware.Auth) CampaignUsecase {
	return &campaignUsecaseImpl{
		campaignRepo: campaignRepo,
		auth: auth,
	}
}