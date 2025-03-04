package handlers

import (
	"bougette-backend/cmd/api/requests"
	"bougette-backend/cmd/api/services"
	"bougette-backend/common"
	"errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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
	return common.SendSuccessResponse(c, "User Registration successful", registeredUser)
}
