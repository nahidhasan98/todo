package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"github.com/nahidhasan98/todo/auth"
	"github.com/nahidhasan98/todo/config"
	"github.com/nahidhasan98/todo/todo"
	"github.com/nahidhasan98/todo/user"
)

func initializeAllServices(router *gin.RouterGroup, dbSession *mgo.Session) {
	authService := auth.Init(router, dbSession)

	user.Init(router, dbSession, authService)
	todo.Init(router, dbSession, authService)
}

func main() {
	router := gin.Default()

	v1 := router.Group("/api")

	session, err := mgo.Dial(config.DbConnectionString)
	if err != nil {
		fmt.Println("Database connection failed!")
		return
	}

	fmt.Println("Server running on port 8080...")
	initializeAllServices(v1, session)

	router.Run(":8080")
}
