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

func (h *handlerStruct) loginHandler(ctx *gin.Context) {
	var reqUser User
	err := ctx.BindJSON(&reqUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &Response{
			Status:  "error",
			Message: fmt.Sprintf("%v", "invalid JSON object"),
		})
		//responseMessage(ctx, "error", http.StatusBadRequest, "invalid JSON object")
		return
	}

	authorizedUser, err := h.authService.Authenticate(&reqUser)
	if err != nil {
		responseMessage(ctx, "error", http.StatusBadRequest, "wrong credentials")
		return
	}

	token, err := h.authService.GenerateToken(authorizedUser)
	if err != nil {
		responseMessage(ctx, "error", http.StatusInternalServerError, "internal server error")
		return
	}

	ctx.JSON(http.StatusOK, token)
}

func responseMessage(ctx *gin.Context, status string, code int, msg interface{}) {
	ctx.JSON(code, &Response{
		Status:  status,
		Message: fmt.Sprintf("%v", msg),
	})
}

func makeHTTPHandlers(router *gin.RouterGroup, authService *AuthService) {
	h := &handlerStruct{
		authService: authService,
	}

	router.POST("/login", h.loginHandler)
}
