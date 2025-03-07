package middlewares

import (
	"bougette-backend/common"
	"bougette-backend/internal/models"
	"errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"strings"
)

type AppMiddlewares struct {
	DB     *gorm.DB
	Logger echo.Logger
}

func (appMiddlewares *AppMiddlewares) AuthenticationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Add("Vary", "Authorization")
		authHeader := c.Request().Header.Get("Authorization")
		if strings.HasPrefix(authHeader, "Bearer ") == false {
			return common.SendUnauthorizationErrorResponse(c, "Please provide a Bearer token")
		}
		authHeaderSplit := strings.Split(authHeader, " ")
		accessToken := authHeaderSplit[1]
		claims, err := common.ParseJWTSignedAccessToken(accessToken)
		if err != nil {
			return common.SendUnauthorizationErrorResponse(c, err.Error())
		}

		if common.IsClaimExpired(claims) {
			return common.SendUnauthorizationErrorResponse(c, "Token is expired")
		}
		var user models.UserModel
		result := appMiddlewares.DB.First(&user, claims.ID)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return common.SendUnauthorizationErrorResponse(c, "invalid access token")
		}
		if result.Error != nil {
			return common.SendUnauthorizationErrorResponse(c, "invalid access token")
		}

		c.Set("user", user)

		return next(c)
	}
}
