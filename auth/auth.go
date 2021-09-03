package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
)

func Init(router *gin.RouterGroup, dbSession *mgo.Session) *AuthService {
	authRepo := NewRepository(dbSession)
	authService := NewAuthService(authRepo)
	makeHTTPHandlers(router, authService)

	return authService
}
