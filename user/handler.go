package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nahidhasan98/todo/auth"
)

type HandlerInterface interface {
	getAllUserHandler(ctx *gin.Context)
	getSingleUserHandler(ctx *gin.Context)
}

type handlerStruct struct {
	userService *UserService
	authService *auth.AuthService
}

func makeHTTPHandlers(router *gin.RouterGroup, userService *UserService, authService *auth.AuthService) {
	h := &handlerStruct{
		userService: userService,
		authService: authService,
	}

	router.GET("user", h.getAllUserHandler)
	router.GET("user/:id", h.getSingleUserHandler)
}

func (handler *handlerStruct) getAllUserHandler(ctx *gin.Context) {
	_, err := handler.authService.ValidateToken(ctx)
	if err != nil { // limited access. response anly user list
		uData, err2 := handler.userService.GetAllUser()
		if err2 != nil {
			ctx.JSON(http.StatusBadRequest, UserResponse{
				UData:   &[]Data{},
				Err:     err2.Error(),
				Message: err.Error() + " | limited access",
			})
			return
		}

		ctx.JSON(http.StatusOK, UserResponse{
			UData:   uData,
			Err:     "",
			Message: err.Error() + " | limited access",
		})
		return
	}

	uData, err := handler.userService.GetAllUserWithTask()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, UserResponse{
			UData: &[]Data{},
			Err:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, UserResponse{
		UData: uData,
		Err:   "",
	})
}

func (handler *handlerStruct) getSingleUserHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	_, err := handler.authService.ValidateToken(ctx)
	if err != nil { // limited access. response anly user list
		user, err2 := handler.userService.GetSingleUser(id)

		if err2 != nil {
			ctx.JSON(http.StatusBadRequest, UserResponse{
				UData:   &[]Data{*user},
				Err:     err2.Error(),
				Message: err.Error() + " | limited access",
			})
			return
		}

		ctx.JSON(http.StatusOK, UserResponse{
			UData:   &[]Data{*user},
			Err:     "",
			Message: err.Error() + " | limited access",
		})
		return
	}

	uData, err := handler.userService.GetSingleUserWithTask(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, UserResponse{
			UData: &[]Data{},
			Err:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, UserResponse{
		UData: &[]Data{*uData},
		Err:   "",
	})
}
