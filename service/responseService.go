package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nahidhasan98/todo/model"
)

func DisplayError(ctx *gin.Context, code int, msg interface{}) {
	ctx.JSON(code, &model.Response{
		Status:  "error",
		Message: fmt.Sprintf("%v", msg),
	})
}

func DisplaySuccess(ctx *gin.Context, code int, msg interface{}) {
	ctx.JSON(code, &model.Response{
		Status:  "success",
		Message: fmt.Sprintf("%v", msg),
	})
}
