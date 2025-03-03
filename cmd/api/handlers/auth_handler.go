package handlers

import (
	"bougette-backend/cmd/api/requests"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) RegisterHandler(c echo.Context) error {
	payload := new(requests.RegisterUserRequest)
	if err := (&echo.DefaultBinder{}).BindBody(c, payload); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	c.Logger().Info(payload)
	validationErrors := h.ValidateBodyRequest(c, *payload)

	if validationErrors != nil {
		return c.JSON(http.StatusBadRequest, validationErrors)
	}

	return c.JSON(http.StatusBadRequest, validationErrors)
}
