package todo

import (
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"github.com/nahidhasan98/todo/auth"
)

func Init(router *gin.RouterGroup, dbSession *mgo.Session, authService *auth.AuthService) {
	todoRepo := NewRepository(dbSession)
	todoService := NewUserService(todoRepo)
	makeHTTPHandlers(router, todoService, authService)
}
