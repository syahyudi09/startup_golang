package handler

import (
	"startup/usecase"

	"github.com/gin-gonic/gin"
)

type CampaignHandler interface {
}

type campaignHandlerImpl struct {
	srv *gin.Engine
	campaign usecase.CampaignUsecase
}

func NewCampaignHandler(srv *gin.Engine, campaign usecase.CampaignUsecase) CampaignHandler {
	Handler := &campaignHandlerImpl{
		campaign: campaign,
	}
	return Handler
}