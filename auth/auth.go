package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
)

func Init(router *gin.RouterGroup, dbSession *mgo.Session) {
	repoService := NewRepository(dbSession)
	authService := NewAuthService(repoService)
	makeHTTPHandlers(router, authService)
}
