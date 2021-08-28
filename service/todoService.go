package service

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/nahidhasan98/todo/db"
	"github.com/nahidhasan98/todo/model"
)

func GetSingleTodo(author, id string) (*model.Todo, error) {
	for _, val := range db.Todo {
		if val.ID == id {
			if val.Author != author {
				return nil, errors.New("not authorized")
			}

			return &val, nil
		}
	}

	return nil, errors.New("no todo found")
}

func GetAllTodo(author string) []model.Todo {
	var todos []model.Todo

	for _, val := range db.Todo {
		if val.Author == author {
			todos = append(todos, val)
		}
	}

	return todos
}

func DeleteSingleTodo(author, id string) (int, error) {
	for key, val := range db.Todo {
		if val.ID == id {
			if val.Author != author {
				return http.StatusUnauthorized, errors.New("not authorized")
			}

			db.Todo = append(db.Todo[:key], db.Todo[key+1:]...)
			return 0, nil
		}
	}

	return http.StatusBadRequest, errors.New("no todo found")
}

func DeleteAllTodo(author string) {
	for key, val := range db.Todo {
		if val.Author == author {
			db.Todo = append(db.Todo[:key], db.Todo[key+1:]...)
			break
		}
	}
}

func AddTodo(todo model.Todo, author string) {
	data := model.Todo{
		ID:      getID(),
		Task:    todo.Task,
		At:      todo.At,
		Message: todo.Message,
		Author:  author,
	}
	db.Todo = append(db.Todo, data)
}

func UpdateTodo(todo model.Todo, id string, author string) (int, error) {
	for key, val := range db.Todo {
		if val.ID == id {
			if val.Author != author {
				return http.StatusUnauthorized, errors.New("not authorized")
			}

			if todo.Task != "" {
				db.Todo[key].Task = todo.Task
			}

			if todo.At != 0 {
				db.Todo[key].At = todo.At
			}

			if todo.Task != "" {
				db.Todo[key].Message = todo.Message
			}

			return 0, nil
		}
	}

	return http.StatusBadRequest, errors.New("no todo found")
}

func getID() string {
	return "t10" + strconv.Itoa(len(db.Todo)+1)
}
