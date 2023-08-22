package handler

import (
	"fmt"
	"net/http"
	"startup/helper"
	"startup/model"
	"startup/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CampaignHandler interface {
}

type campaignHandlerImpl struct {
	srv *gin.Engine
	campaignUsecae usecase.CampaignUsecase
}

func (ch *campaignHandlerImpl) GetCampaign(ctx *gin.Context) {
	userIDStr := ctx.Query("user_id")
	userID, _ := strconv.Atoi(userIDStr)

	var campaigns []*model.CampaignModel
	var err error

	if userID != 0 {
		campaign, err := ch.campaignUsecae.FindByID(userID)
		fmt.Println("err", err)
		if err != nil {
			response := helper.APIResponse(
				"Failed to get campaign by ID", 
				http.StatusBadRequest, 
				"error", 
				nil,
			)
			ctx.JSON(http.StatusBadRequest, response)
			return
		}
		campaigns = append(campaigns, campaign)
	} else {
		campaigns, err = ch.campaignUsecae.FindAll()
		fmt.Println("err", err)
		if err != nil {
			response := helper.APIResponse(
				"Failed to get all campaigns", 
				http.StatusUnprocessableEntity, 
				"error", 
				nil,
			)
			ctx.JSON(http.StatusUnprocessableEntity, response)
			return
		}
	}

	response := helper.APIResponse("Successfully retrieved campaigns", http.StatusOK, "success", campaigns)
	ctx.JSON(http.StatusOK, response)
}


func NewCampaignHandler(srv *gin.Engine, campaign usecase.CampaignUsecase) CampaignHandler {
	Handler := &campaignHandlerImpl{
		campaignUsecae: campaign,
		srv: srv,
	}

	srv.GET("/campagins", Handler.GetCampaign)
	return Handler
}