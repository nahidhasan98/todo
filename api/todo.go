package api

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/nahidhasan98/todo/model"
	"github.com/nahidhasan98/todo/service"
)

func GetSingleTodo(ctx *gin.Context, claims jwt.MapClaims) {
	author := fmt.Sprintf("%v", claims["username"])
	todoID := ctx.Param("id")

	todo, err := service.GetSingleTodo(author, todoID)
	if err != nil {
		service.DisplayError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, todo)
}

func GetAllTodo(ctx *gin.Context, claims jwt.MapClaims) {
	author := fmt.Sprintf("%v", claims["username"])
	ctx.JSON(http.StatusOK, service.GetAllTodo(author))
}

func DeleteSingleTodo(ctx *gin.Context, claims jwt.MapClaims) {
	author := fmt.Sprintf("%v", claims["username"])
	todoID := ctx.Param("id")

	statusCode, err := service.DeleteSingleTodo(author, todoID)
	if err != nil {
		service.DisplayError(ctx, statusCode, err)
		return
	}

	service.DisplaySuccess(ctx, http.StatusOK, "data successfully deleted")
}

func DeleteAllTodo(ctx *gin.Context, claims jwt.MapClaims) {
	author := fmt.Sprintf("%v", claims["username"])
	service.DeleteAllTodo(author)

	service.DisplaySuccess(ctx, http.StatusOK, "todo list reset successfully")
}

func AddTodo(ctx *gin.Context, claims jwt.MapClaims) {
	var todo model.Todo
	err := ctx.BindJSON(&todo)
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

func UpdateTodo(ctx *gin.Context, claims jwt.MapClaims) {
	var todo model.Todo
	err := ctx.BindJSON(&todo)
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

	statusCode, err := service.UpdateTodo(todo, todoID, author)
	if err != nil {
		service.DisplayError(ctx, statusCode, err)
		return
	}

	service.DisplaySuccess(ctx, http.StatusOK, "data successfully updated")
}
