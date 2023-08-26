package handler

import (
	"fmt"
	"net/http"
	"startup/formatter"
	"startup/helper"
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

// func (ch *campaignHandlerImpl) GetCampaignID(ctx *gin.Context) {
// 	userIDStr := ctx.Param("user_id")
// 	if userIDStr == ""{
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"errorMessage": "Id tidak boleh kosong",
// 		})
// 		return
// 	}

// 	id, err := strconv.Atoi(userIDStr)
// 	if err != nil{
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"errorMessage": "user_id harus angka",
// 		})
// 		return
// 	}

// 	campaigns, err := ch.campaignUsecae.FindByID(id)
// 	if err != nil{
// 		fmt.Printf("error an ch.campaignUsecae.FindByID: %v", err)
// 		response := helper.APIResponse("Get Campaign Failed", http.StatusBadRequest, "error", nil)
// 		ctx.JSON(http.StatusInternalServerError, response)
// 		return
// 	}

// 	formatter := formatter.FormatterCampaign(campaigns)
// 	response := helper.APIResponse("Successfully retrieved campaigns", http.StatusOK, "success", formatter)
// 	ctx.JSON(http.StatusOK, response)
// }

// func (campaignHandler *campaignHandlerImpl) GetCampaignByID(ctx *gin.Context) {
// 	userIDStr := ctx.Param("user_id")
// 	if userIDStr == "" {
// 		response := helper.APIResponse("user_id not valid", http.StatusInternalServerError, "error", nil)
// 		ctx.JSON(http.StatusInternalServerError, response)
// 		return
// 	}

// 	id, err := strconv.Atoi(userIDStr)
// 	if err != nil {
// 		fmt.Printf("error in campaignHandler.campaignUsecae.FindByID: %v", err)
// 		response := helper.APIResponse("user_id not valid", http.StatusInternalServerError, "error", nil)
// 		ctx.JSON(http.StatusInternalServerError, response)
// 		return
// 	}

// 	campaign, err := campaignHandler.campaignUsecae.FindByID(id)
// 	if err != nil {
// 		fmt.Printf("error in campaignHandler.campaignUsecae.FindByID: %v", err)
// 		response := helper.APIResponse("Gagal mendapatkan kampanye", http.StatusInternalServerError, "error", nil)
// 		ctx.JSON(http.StatusInternalServerError, response)
// 		return
// 	}
	
// 	formatter := []formatter.CampaignFormatter{}
// 	if campaign == nil {
// 		formatter = formatter.FormatterCampaign(campaign)
// 		response := helper.APIResponse("Kampanye tidak ditemukan", http.StatusNotFound, "error", errors)
// 		ctx.JSON(http.StatusNotFound, response)
// 		return
// 	}

// 	response := helper.APIResponse("Berhasil mendapatkan kampanye", http.StatusOK, "success", formatter)
// 	ctx.JSON(http.StatusOK, response)
	
// }

func (campaignHandler *campaignHandlerImpl) GetCampaignByID(ctx *gin.Context) {
	userIDStr := ctx.Param("user_id")
	if userIDStr == "" {
		response := helper.APIResponse("user_id not valid", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	id, err := strconv.Atoi(userIDStr)
	if err != nil {
		fmt.Printf("error in campaignHandler.campaignUsecae.FindByID: %v", err)
		response := helper.APIResponse("user_id not valid", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	campaign, err := campaignHandler.campaignUsecae.FindByID(id)
	if err != nil {
		fmt.Printf("error in campaignHandler.campaignUsecae.FindByID: %v", err)
		response := helper.APIResponse("Gagal mendapatkan kampanye", http.StatusInternalServerError, "error", nil)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	
	if campaign == nil {
		response := helper.APIResponse("Kampanye tidak ditemukan", http.StatusNotFound, "error", nil)
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	formatter := formatter.FormatCampaign(campaign)
	response := helper.APIResponse("Berhasil mendapatkan kampanye", http.StatusOK, "success", formatter)
	ctx.JSON(http.StatusOK, response)
}


func (ch *campaignHandlerImpl) GetCampaignAll(ctx *gin.Context){
	campaigns, err := ch.campaignUsecae.FindAll()
	if err != nil{
		fmt.Printf("error an ch.campaignUsecae.FindByID: %v", err)
		response := helper.APIResponse("Get Campaign Failed", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	formatter := formatter.FormatCampaigns(campaigns)
	response := helper.APIResponse("Successfully retrieved campaigns", http.StatusOK, "success", formatter)
	ctx.JSON(http.StatusOK, response)
}

func NewCampaignHandler(srv *gin.Engine, campaign usecase.CampaignUsecase) CampaignHandler {
	Handler := &campaignHandlerImpl{
		campaignUsecae: campaign,
		srv: srv,
	}



	srv.GET("/campagins/:user_id", Handler.GetCampaignByID)
	srv.GET("/campagins", Handler.GetCampaignAll)
	return Handler
}