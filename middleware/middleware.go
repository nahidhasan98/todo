package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nahidhasan98/todo/auth"
)

func Authorization(auth *auth.AuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, err := auth.ParseToken(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"Err": err.Error(),
			})
		}

		ctx.Next()
	}
}
