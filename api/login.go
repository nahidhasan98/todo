package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nahidhasan98/todo/model"
	"github.com/nahidhasan98/todo/service"
)

func Login(ctx *gin.Context) {
	var reqUser model.Credentials
	err := ctx.BindJSON(&reqUser)
	if err != nil {
		service.DisplayError(ctx, http.StatusBadRequest, "invalid JSON object")
		return
	}

	authorizedUser, err := service.Authenticate(&reqUser)
	if err != nil {
		service.DisplayError(ctx, http.StatusUnauthorized, "wrong credentials")
		return
	}

	token, err := service.GenerateToken(authorizedUser)
	if err != nil {
		service.DisplayError(ctx, http.StatusInternalServerError, "internal server error")
		return
	}

	ctx.JSON(http.StatusOK, token)
}
