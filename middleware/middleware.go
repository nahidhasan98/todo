package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/nahidhasan98/todo/service"
)

func GetPass(sendToAPI func(ctx *gin.Context, claims jwt.MapClaims)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, err := service.ValidateToken(ctx)
		if err != nil {
			service.DisplayError(ctx, http.StatusUnauthorized, err)
			return
		}

		sendToAPI(ctx, claims)
	}
}

func AutoPass(sendToAPI func(ctx *gin.Context)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sendToAPI(ctx)
	}
}

func SpecialPass(sendToAPI func(ctx *gin.Context, err error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, err := service.ValidateToken(ctx)

		sendToAPI(ctx, err)
	}
}
