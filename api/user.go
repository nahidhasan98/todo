package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nahidhasan98/todo/model"
	"github.com/nahidhasan98/todo/service"
)

func GetUsers(ctx *gin.Context) {
	_, err := service.ValidateToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, service.GetAllUser(err, false))
		return
	}

	ctx.JSON(http.StatusOK, service.GetAllUser(nil, true))
}

func GetUser(ctx *gin.Context) {
	userID := ctx.Param("id")
	userData := model.Response{}
	var err error

	_, err = service.ValidateToken(ctx)
	if err != nil {
		userData, err = service.GetSingleUser(userID, err, false)
	} else {
		userData, err = service.GetSingleUser(userID, nil, true)
	}

	if err != nil {
		service.DisplayError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, userData)
}
