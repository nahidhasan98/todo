package service

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/nahidhasan98/todo/db"
	"github.com/nahidhasan98/todo/model"
)

func GetTodo(author, id string) (*model.Todo, error) {
	for _, val := range db.Todo {
		if val.Author == author && val.ID == id {
			return &val, nil
		}
	}

	return nil, errors.New("no todo found")
}

func GetTodos(author string) []model.Todo {
	var todos []model.Todo

	for _, val := range db.Todo {
		if val.Author == author {
			todos = append(todos, val)
		}
	}

	return todos
}

func DeleteTodo(author, id string) error {
	var todos []model.Todo
	flag := false

	for _, val := range db.Todo {
		if val.Author == author && val.ID == id {
			flag = true
		} else {
			todos = append(todos, val)
		}
	}

	db.Todo = todos

	if flag {
		return nil
	} else {
		return errors.New("no todo found")
	}
}

func DeleteTodos(author string) {
	var todos []model.Todo

	for _, val := range db.Todo {
		if val.Author != author {
			todos = append(todos, val)
		}
	}

	db.Todo = todos
}

func AddTodo(todo model.Todo, claims jwt.MapClaims) {
	data := model.Todo{
		ID:      getID(),
		Task:    todo.Task,
		At:      todo.At,
		Message: todo.Message,
		Author:  fmt.Sprintf("%v", claims["username"]),
	}
	db.Todo = append(db.Todo, data)
}

func UpdateTodo(todo model.Todo, claims jwt.MapClaims) {
	author := claims["username"]

	for key, val := range db.Todo {
		if val.Author == author {
			if todo.Task != "" {
				db.Todo[key].Task = todo.Task
			}

			if todo.At != 0 {
				db.Todo[key].At = todo.At
			}

			if todo.Task != "" {
				db.Todo[key].Message = todo.Message
			}
		}
	}
}

func getID() string {
	return "t10" + strconv.Itoa(len(db.Todo)+1)
}
