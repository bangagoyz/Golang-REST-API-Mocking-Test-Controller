package middlewares

import (
	"chapter3_2/helper"
	"chapter3_2/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware(ctx *gin.Context) {
	auth := ctx.GetHeader("Authorization")

	if auth == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.FailedResponse{
			Response: models.Meta{
				Code:    http.StatusUnauthorized,
				Message: "Insert Access Token",
			},
			Error: models.ErrorNotAuthorized.Err,
		})
	}

	token := strings.Split(auth, " ")[1]

	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.FailedResponse{
			Response: models.Meta{
				Code:    http.StatusUnauthorized,
				Message: http.StatusText(http.StatusUnauthorized),
			},
			Error: models.ErrorNotAuthorized.Err,
		})
	}

	jwtToken, err := helper.VerifyAccessToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.FailedResponse{
			Response: models.Meta{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			},
			Error: err.Error(),
		})
		return
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.FailedResponse{
			Response: models.Meta{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			},
			Error: err.Error(),
		})
		return

	}

	ctx.Set("user_id", claims["user_id"])

	ctx.Next()
}
