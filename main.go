package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nahidhasan98/todo/api"
)

func main() {
	router := gin.Default()

	// login group
	router.POST("/api/login", func(ctx *gin.Context) { api.Login(ctx) })

	// user group
	user := router.Group("/api/user")
	{
		user.GET("", func(ctx *gin.Context) { api.GetUsers(ctx) })
		user.GET("/:id", func(ctx *gin.Context) { api.GetUser(ctx) })
	}

	// todo group
	todo := router.Group("/api/todo")
	{
		todo.GET("/:id", func(ctx *gin.Context) { api.GetTodo(ctx) })
		todo.GET("", func(ctx *gin.Context) { api.GetTodos(ctx) })

		todo.DELETE("/:id", func(ctx *gin.Context) { api.DeleteTodo(ctx) })
		todo.DELETE("", func(ctx *gin.Context) { api.DeleteTodos(ctx) })

		todo.POST("", func(ctx *gin.Context) { api.AddTodo(ctx) })
		todo.PATCH("/:id", func(ctx *gin.Context) { api.UpdateTodo(ctx) })
	}

	router.Run(":8080")
}
