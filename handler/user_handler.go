package handler

import (
	"fmt"
	"net/http"
	"startup/helper"
	"startup/model"
	"startup/usecase"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
}

type userhandlerImpl struct {
	userUsecase usecase.Userusecase
	srv *gin.Engine
}

func (userHandler *userhandlerImpl) RegisterUser(ctx *gin.Context){
	register := &model.RegisterUserInput{}
	err := ctx.ShouldBindJSON(&register)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Registrasi Account Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity,response)
		return
	}

	err = userHandler.userUsecase.RegisterUser(register)
	if err != nil{
		fmt.Printf("error an  userHandler.userUsecase.RegisterUser: %v", err)
		response := helper.APIResponse("Registrasi Account Failed", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	// response dari helper
	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", register)
	ctx.JSON(http.StatusOK, response)
}

func (userHandler *userhandlerImpl) LoginUser(ctx *gin.Context) {
	input := &model.LoginUser{}
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity,response)
		return
	}

	err = userHandler.userUsecase.LoginUser(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity,response)
		return
	}

	response := helper.APIResponse("Successfuly Login", http.StatusOK, "success", input.Password)
	ctx.JSON(http.StatusOK, response)
}

func (UserHandler *userhandlerImpl) CheckEmailAvalible(ctx *gin.Context) {
	var input model.CheckEmailAvailable
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Email checking Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity,response)
		return
	}

	isEmailAvailible, err := UserHandler.userUsecase.IsAvailableEmail(&input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server Error"}
		response := helper.APIResponse("Email checking Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity,response)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailible,
	}

	metaMessage := "Email Telah Terdaftar"

	if isEmailAvailible {
		metaMessage = "Email Tersedia"
	}
	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	ctx.JSON(http.StatusOK,response)
}



func NewUserHandler(srv *gin.Engine,user usecase.Userusecase) UserHandler{
	Handler := userhandlerImpl{
		userUsecase: user,
		srv: srv,
	}	

	srv.POST("/register", Handler.RegisterUser)
	srv.POST("/login", Handler.LoginUser)
	srv.POST("/email_chekers", Handler.CheckEmailAvalible)
	

	return Handler
}