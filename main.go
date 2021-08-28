package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nahidhasan98/todo/api"
	"github.com/nahidhasan98/todo/middleware"
)

func main() {
	router := gin.Default()

	// login group
	router.POST("/api/login", middleware.AutoPass(api.Login))

	// user group
	user := router.Group("/api/user")
	{
		user.GET("", middleware.SpecialPass(api.GetUserData))
		user.GET("/:id", middleware.SpecialPass(api.GetUserData))
	}

	// todo group
	todo := router.Group("/api/todo")
	{
		todo.GET("/:id", middleware.GetPass(api.GetSingleTodo))
		todo.GET("", middleware.GetPass(api.GetAllTodo))

		todo.DELETE("/:id", middleware.GetPass(api.DeleteSingleTodo))
		todo.DELETE("", middleware.GetPass(api.DeleteAllTodo))

		todo.POST("", middleware.GetPass(api.AddTodo))
		todo.PATCH("/:id", middleware.GetPass(api.UpdateTodo))
	}

	router.Run(":8080")
}
