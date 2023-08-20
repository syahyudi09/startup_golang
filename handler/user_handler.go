package handler

import (
	"fmt"
	"net/http"
	"startup/helper"
	"startup/middleware"
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
		var input model.LoginUser
		err := ctx.ShouldBindJSON(&input)
		if err != nil {
			errors := helper.FormatValidationError(err)
			errorMessage := gin.H{"errors": errors}
			response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
			ctx.JSON(http.StatusUnprocessableEntity,response)
			return
		}

		token, err := userHandler.userUsecase.LoginUser(input)
		if err != nil {
			errorMessage := gin.H{"errors": err.Error()}
			response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
			ctx.JSON(http.StatusUnprocessableEntity,response)
			return
		}

		data := gin.H{
			"token": token,}

		response := helper.APIResponse("Successfuly Login", http.StatusOK, "success", data)
		ctx.JSON(http.StatusOK, response)
	}

	func (UserHandler *userhandlerImpl) CheckEmailAvailable(ctx *gin.Context) {
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
			errorMessage := gin.H{
				"errors": "Server Error",
			}
			response := helper.APIResponse(
				"Email checking Failed", 
				http.StatusUnprocessableEntity, 
				"error", 
				errorMessage,
			)
			ctx.JSON(http.StatusUnprocessableEntity,response)
			return
		}

		data := gin.H{
			"is_available": isEmailAvailible,
		}

		var metaMessage string
		if isEmailAvailible {
			metaMessage = "Email Telah Terdaftar"
		}else{
			metaMessage = "Email Tersedia"
		}
		response := helper.APIResponse(
			metaMessage, 
			http.StatusOK, 
			"success", 
			data,	)
		ctx.JSON(http.StatusOK,response)
	}

	func (UserHandler *userhandlerImpl) UploadAvatar(ctx *gin.Context) {
		file, err := ctx.FormFile("avatar")
		if err != nil {
			response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", nil)
			ctx.JSON(http.StatusBadRequest, response)
			return
		}

		// Pastikan file tidak bernilai nil sebelum melanjutkan
		if file == nil {
			response := helper.APIResponse("No file uploaded", http.StatusBadRequest, "error", nil)
			ctx.JSON(http.StatusBadRequest, response)
			return
		}

		// Lanjutkan dengan pemrosesan file
		path := "images/" + file.Filename

		err = ctx.SaveUploadedFile(file, path)
		if err != nil {
			response := helper.APIResponse("Failed to upload avatar image", http.StatusInternalServerError, "error", nil)
			ctx.JSON(http.StatusInternalServerError, response)
			return
		}

		// Mendapatkan ID pengguna dari autentikasi, misalnya menggunakan JWT token
		userID := 2

		err = UserHandler.userUsecase.UpdateAvatar(userID, path)
		if err != nil {
			response := helper.APIResponse("Failed to update avatar", http.StatusInternalServerError, "error", nil)
			ctx.JSON(http.StatusInternalServerError, response)
			return
		}

		data := gin.H{
			"is_Uploaded": true,
		}

		response := helper.APIResponse("Avatar successfully uploaded and updated", http.StatusOK, "success", data)
		ctx.JSON(http.StatusOK, response)
	}

func NewUserHandler(srv *gin.Engine, user usecase.Userusecase) UserHandler {
	handler := &userhandlerImpl{
		userUsecase:   user,
			srv:           srv,
	}

	auth := middleware.NewJwtService()
	middleware := NewMiddleware(auth)

	authenticated := srv.Group("/")
	authenticated.Use(middleware.AuthMiddleware())
	
	srv.POST("/register", handler.RegisterUser)
	srv.POST("/login", handler.LoginUser)
	srv.POST("/email_checkers", handler.CheckEmailAvailable)
	authenticated.POST("/avatars", handler.UploadAvatar)
	
		return handler
	}