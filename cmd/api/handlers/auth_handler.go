package handlers

import (
	"bougette-backend/cmd/api/requests"
	"bougette-backend/cmd/api/services"
	"bougette-backend/common"
	"bougette-backend/internal/mailer"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"os"
)

func (h *Handler) RegisterHandler(c echo.Context) error {
	payload := new(requests.RegisterUserRequest)
	if err := (&echo.DefaultBinder{}).BindBody(c, payload); err != nil {
		return common.SendBadRequestResponse(c, err.Error())
	}
	c.Logger().Info(payload)
	validationErrors := h.ValidateBodyRequest(c, *payload)

	if validationErrors != nil {
		return common.SendFailedValidationResponse(c, validationErrors)
	}

	userService := services.NewUserService(h.DB)

	_, err := userService.GetUserByEmail(payload.Email)
	if errors.Is(err, gorm.ErrRecordNotFound) == false {
		common.SendBadRequestResponse(c, "Email has already been taken")
	}
	registeredUser, err := userService.RegisterUser(payload)
	if err != nil {
		common.SendInternalServerErrorResponse(c, err.Error())
	}
	mailData := mailer.EmailData{
		Subject: "Welcome To " + os.Getenv("APP_NAME") + " Signup",
		Meta: struct {
			FirstName string
			LoginLink string
		}{
			FirstName: *registeredUser.FirstName,
			LoginLink: "#",
		},
	}
	// Send a welcome message to the user
	err = h.Mailer.Send(payload.Email, "hello.html", mailData)
	if err != nil {
		h.Logger.Error(err)
	}
	return common.SendSuccessResponse(c, "User Registration successful", registeredUser)
}

func (h *Handler) LoginHandler(c echo.Context) error {
	userService := services.NewUserService(h.DB)

	payload := new(requests.LoginUserRequest)
	if err := (&echo.DefaultBinder{}).BindBody(c, payload); err != nil {
		return common.SendBadRequestResponse(c, err.Error())
	}
	c.Logger().Info(payload)
	validationErrors := h.ValidateBodyRequest(c, *payload)

	if validationErrors != nil {
		return common.SendFailedValidationResponse(c, validationErrors)
	}

	userRetrieved, err := userService.GetUserByEmail(payload.Email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return common.SendBadRequestResponse(c, "Invalid credential")
	}
	fmt.Println(*userRetrieved)
	if common.ComparePasswordHash(payload.Password, userRetrieved.Password) == false {
		return common.SendBadRequestResponse(c, "Invalid credential")
	}
	// send response with token
	accessToken, refreshToken, err := common.GenerateJWT(*userRetrieved)
	if err != nil {
		return common.SendInternalServerErrorResponse(c, err.Error())
	}
	return common.SendSuccessResponse(c, "User Login", map[string]interface{}{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
		"user":         userRetrieved,
	})

}
