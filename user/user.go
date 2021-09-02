package user

import (
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"github.com/nahidhasan98/todo/auth"
)

func Init(router *gin.RouterGroup, dbSession *mgo.Session) {
	authRepo := auth.NewRepository(dbSession)
	authService := auth.NewAuthService(authRepo)

	userRepo := NewRepository(dbSession)
	userService := NewUserService(userRepo)
	makeHTTPHandlers(router, userService, authService)
}
