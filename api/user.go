package api

import (
	"github.com/gin-gonic/gin"
	"github.com/nahidhasan98/todo/service"
)

func GetUserData(ctx *gin.Context, err error) {
	statusCode, data := service.GetUserData(ctx, err)

	ctx.JSON(statusCode, data)
}
