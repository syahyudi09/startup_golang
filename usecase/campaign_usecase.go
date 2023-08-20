package usecase

import (
	"startup/middleware"
	"startup/repo"
)

type CampaignUsecase interface {
}

type campaignUsecaseImpl struct {
	campaignRepo repo.CampaignRepo
	auth middleware.Auth
}

func NewCampaignUsecase(campaignRepo repo.CampaignRepo, auth middleware.Auth) CampaignUsecase {
	return &campaignUsecaseImpl{
		campaignRepo: campaignRepo,
		auth: auth,
	}
}