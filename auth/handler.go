package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerInterface interface {
	loginHandler(ctx *gin.Context)
}

type handlerStruct struct {
	authService *AuthService
}

func makeHTTPHandlers(router *gin.RouterGroup, authService *AuthService) {
	handler := &handlerStruct{
		authService: authService,
	}

	router.POST("/login", handler.loginHandler)
}

func (handler *handlerStruct) loginHandler(ctx *gin.Context) {
	var reqUser User
	err := ctx.BindJSON(&reqUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &Response{
			Status:  "error",
			Message: fmt.Sprintf("%v", "invalid JSON object"),
		})
		return
	}

	authorizedUser, err := handler.authService.Authenticate(&reqUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &Response{
			Status:  "error",
			Message: fmt.Sprintf("%v", "wrong credentials"),
		})
		return
	}

	token, err := handler.authService.GenerateToken(authorizedUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &Response{
			Status:  "error",
			Message: fmt.Sprintf("%v", "internal server error"),
		})
		return
	}

	ctx.JSON(http.StatusOK, token)
}
