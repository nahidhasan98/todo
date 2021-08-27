package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nahidhasan98/todo/model"
	"github.com/nahidhasan98/todo/service"
)

func GetTodo(ctx *gin.Context) {
	claims, err := service.ValidateToken(ctx)
	if err != nil {
		service.DisplayError(ctx, http.StatusUnauthorized, err)
		return
	}

	author := fmt.Sprintf("%v", claims["username"])
	todoID := ctx.Param("id")
	todo, err := service.GetTodo(author, todoID)
	if err != nil {
		service.DisplayError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, todo)
}

func GetTodos(ctx *gin.Context) {
	claims, err := service.ValidateToken(ctx)
	if err != nil {
		service.DisplayError(ctx, http.StatusUnauthorized, err)
		return
	}

	author := fmt.Sprintf("%v", claims["username"])
	ctx.JSON(http.StatusOK, service.GetTodos(author))
}

func DeleteTodo(ctx *gin.Context) {
	claims, err := service.ValidateToken(ctx)
	if err != nil {
		service.DisplayError(ctx, http.StatusUnauthorized, err)
		return
	}

	author := fmt.Sprintf("%v", claims["username"])
	todoID := ctx.Param("id")

	err = service.DeleteTodo(author, todoID)
	if err != nil {
		service.DisplayError(ctx, http.StatusBadRequest, err)
		return
	}

	service.DisplaySuccess(ctx, http.StatusOK, "data successfully deleted")
}

func DeleteTodos(ctx *gin.Context) {
	claims, err := service.ValidateToken(ctx)
	if err != nil {
		service.DisplayError(ctx, http.StatusUnauthorized, err)
		return
	}

	author := fmt.Sprintf("%v", claims["username"])
	service.DeleteTodos(author)

	service.DisplaySuccess(ctx, http.StatusOK, "todo list reset successfully")
}

func AddTodo(ctx *gin.Context) {
	claims, err := service.ValidateToken(ctx)
	if err != nil {
		service.DisplayError(ctx, http.StatusUnauthorized, err)
		return
	}

	var todo model.Todo
	err = ctx.BindJSON(&todo)
	if err != nil {
		service.DisplayError(ctx, http.StatusBadRequest, "invalid JSON object")
		return
	}

	if todo.Task == "" {
		service.DisplayError(ctx, http.StatusBadRequest, "no task provided")
		return
	}

	author := fmt.Sprintf("%v", claims["username"])
	service.AddTodo(todo, author)

	service.DisplaySuccess(ctx, http.StatusOK, "data successfully added")
}

func UpdateTodo(ctx *gin.Context) {
	claims, err := service.ValidateToken(ctx)
	if err != nil {
		service.DisplayError(ctx, http.StatusUnauthorized, err)
		return
	}

	var todo model.Todo
	err = ctx.BindJSON(&todo)
	if err != nil {
		service.DisplayError(ctx, http.StatusBadRequest, "invalid JSON object")
		return
	}

	if todo.Task == "" {
		service.DisplayError(ctx, http.StatusBadRequest, "update failed. no task provided")
		return
	}

	author := fmt.Sprintf("%v", claims["username"])
	todoID := ctx.Param("id")

	err = service.UpdateTodo(todo, todoID, author)
	if err != nil {
		service.DisplayError(ctx, http.StatusBadRequest, err)
		return
	}

	service.DisplaySuccess(ctx, http.StatusOK, "data successfully updated")
}
