package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nahidhasan98/todo/service"
)

func GetUsers(ctx *gin.Context) {
	_, err := service.ValidateToken(ctx)
	if err != nil {
		service.DisplayError(ctx, http.StatusUnauthorized, err)
		return
	}

	ctx.JSON(http.StatusOK, service.GetAllUser())
}

func GetUser(ctx *gin.Context) {
	_, err := service.ValidateToken(ctx)
	if err != nil {
		service.DisplayError(ctx, http.StatusUnauthorized, err)
		return
	}

	userID := ctx.Param("id")

	user, err := service.GetSingleUser(userID)
	if err != nil {
		service.DisplayError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, user)
}
