package user

import (
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
)

func Init(router *gin.RouterGroup, dbSession *mgo.Session) {
	userRepo := NewRepository(dbSession)
	userService := NewUserService(userRepo)
	makeHTTPHandlers(router, userService)
}
