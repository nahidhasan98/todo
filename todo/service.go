package todo

import (
	"fmt"
)

type TodoServiceInterface interface {
	CreateTodo(todo *Todo, author string) error
	GetAllTodo(author string) (*[]Todo, error)
	GetSingleTodo()
	DeleteAllTodo()
	DeleteSingleTodo()
	UpdateTodo()
}

type TodoService struct {
	repoService *repoStruct
}

func (todoService *TodoService) CreateTodo(todo *Todo, author string) error {
	todoPro := &Todo{
		ID:      todoService.repoService.getID(),
		Task:    todo.Task,
		At:      todo.At,
		Message: todo.Message,
		Author:  author,
	}
	err := todoService.repoService.createTodo(todoPro, author)

	return err
}

func (todoService *TodoService) GetAllTodo(author string) (*[]Todo, error) {
	todo, err := todoService.repoService.getAllTodo(author)
	if err != nil {
		return nil, err
	}
	fmt.Println(todo)
	return todo, nil
}

func (todoService *TodoService) GetSingleTodo(author, todoID string) (*Todo, error) {
	todo, err := todoService.repoService.getSingleTodo(author, todoID)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (todoService *TodoService) DeleteAllTodo(author string) error {
	err := todoService.repoService.deleteAllTodo(author)
	if err != nil {
		return err
	}
	return nil
}

func (todoService *TodoService) DeleteSingleTodo(author, todoID string) error {
	err := todoService.repoService.deleteSingleTodo(author, todoID)
	if err != nil {
		return err
	}
	return nil
}

func (todoService *TodoService) UpdateTodo(todo *Todo, author, todoID string) error {
	var todoPro Todo
	if todo.Task != "" {
		todoPro.Task = todo.Task
	}

	if todo.At != 0 {
		todoPro.At = todo.At
	}

	if todo.Task != "" {
		todoPro.Message = todo.Message
	}

	err := todoService.repoService.updateTodo(&todoPro, author, todoID)
	if err != nil {
		return err
	}
	return nil
}

func NewUserService(repo *repoStruct) *TodoService {
	return &TodoService{
		repoService: repo,
	}
}
