package todo

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nahidhasan98/todo/auth"
	"github.com/nahidhasan98/todo/middleware"
)

type HandlerInterface interface {
	createTodoHandler(ctx *gin.Context)
	getAllTodoHandler(ctx *gin.Context)
	getSingleTodoHandler(ctx *gin.Context)
	deleteAllTodoHandler(ctx *gin.Context)
	deleteSingleTodoHandler(ctx *gin.Context)
	updateTodoHandler(ctx *gin.Context)
}

type handlerStruct struct {
	todoService *TodoService
	authService *auth.AuthService
}

func makeHTTPHandlers(router *gin.RouterGroup, todoService *TodoService, authService *auth.AuthService) {
	h := &handlerStruct{
		authService: authService,
		todoService: todoService,
	}

	router.Use(middleware.Authorization(authService))

	router.POST("todo", h.createTodoHandler)
	router.GET("todo", h.getAllTodoHandler)
	router.GET("todo/:id", h.getSingleTodoHandler)
	router.DELETE("todo", h.deleteAllTodoHandler)
	router.DELETE("todo/:id", h.deleteSingleTodoHandler)
	router.PATCH("todo/:id", h.updateTodoHandler)
}

func (handler *handlerStruct) createTodoHandler(ctx *gin.Context) {
	var todo Todo
	err := ctx.BindJSON(&todo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &Response{
			Status:  "error",
			Message: "invalid JSON object",
		})
		return
	}

	if todo.Task == "" {
		ctx.JSON(http.StatusBadRequest, &Response{
			Status:  "error",
			Message: "no task provided",
		})
		return
	}

	claims, _ := handler.authService.ParseToken(ctx)
	author := fmt.Sprintf("%v", claims["username"])

	err = handler.todoService.CreateTodo(&todo, author)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &Response{
			Status:  "error",
			Message: "couldn't store data to database. error: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &Response{
		Status:  "success",
		Message: "data successfully added",
	})
}

func (handler *handlerStruct) getAllTodoHandler(ctx *gin.Context) {
	claims, _ := handler.authService.ParseToken(ctx)
	author := fmt.Sprintf("%v", claims["username"])

	todo, err := handler.todoService.GetAllTodo(author)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &Response{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, todo)
}

func (handler *handlerStruct) getSingleTodoHandler(ctx *gin.Context) {
	claims, _ := handler.authService.ParseToken(ctx)
	author := fmt.Sprintf("%v", claims["username"])
	todoID := ctx.Param("id")

	todo, err := handler.todoService.GetSingleTodo(author, todoID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &Response{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, todo)
}
func (handler *handlerStruct) deleteAllTodoHandler(ctx *gin.Context) {
	claims, _ := handler.authService.ParseToken(ctx)
	author := fmt.Sprintf("%v", claims["username"])

	err := handler.todoService.DeleteAllTodo(author)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &Response{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &Response{
		Status:  "success",
		Message: "data successfully deleted",
	})
}
func (handler *handlerStruct) deleteSingleTodoHandler(ctx *gin.Context) {
	claims, _ := handler.authService.ParseToken(ctx)
	author := fmt.Sprintf("%v", claims["username"])
	todoID := ctx.Param("id")

	err := handler.todoService.DeleteSingleTodo(author, todoID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &Response{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &Response{
		Status:  "success",
		Message: "data successfully deleted",
	})
}
func (handler *handlerStruct) updateTodoHandler(ctx *gin.Context) {
	var todo Todo
	err := ctx.BindJSON(&todo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &Response{
			Status:  "error",
			Message: "invalid JSON object",
		})
		return
	}

	if todo.Task == "" {
		ctx.JSON(http.StatusBadRequest, &Response{
			Status:  "error",
			Message: "no task provided",
		})
		return
	}

	claims, _ := handler.authService.ParseToken(ctx)
	author := fmt.Sprintf("%v", claims["username"])
	todoID := ctx.Param("id")

	err = handler.todoService.UpdateTodo(&todo, author, todoID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &Response{
			Status:  "error",
			Message: "couldn't update data. error: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &Response{
		Status:  "success",
		Message: "data successfully updated",
	})
}
