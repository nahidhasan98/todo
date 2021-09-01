package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerInterface interface {
	getAllUserHandler(ctx *gin.Context)
	getSingleUserHandler(ctx *gin.Context)
}

type handlerStruct struct {
	userService *UserService
}

func makeHTTPHandlers(router *gin.RouterGroup, userService *UserService) {
	h := &handlerStruct{
		userService: userService,
	}

	router.GET("user", h.getAllUserHandler)
	router.GET("user/:id", h.getSingleUserHandler)
}

type getAllUserResponse struct {
	User *[]User `json:"user"`
	Err  string  `json:"err"`
}

func (h *handlerStruct) getAllUserHandler(ctx *gin.Context) {
	user, err := h.userService.GetAllUser()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, getAllUserResponse{
			User: &[]User{},
			Err:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, getAllUserResponse{
		User: user,
		Err:  "",
	})
}

type getSingleUserResponse struct {
	User *User  `json:"user"`
	Err  string `json:"err"`
}

func (h *handlerStruct) getSingleUserHandler(ctx *gin.Context) {
	username := ctx.Param("id")
	user, err := h.userService.GetSingleUser(username)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, getSingleUserResponse{
			User: &User{},
			Err:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, getSingleUserResponse{
		User: user,
		Err:  "",
	})
}
