package service

import (
	"errors"
	"strconv"

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
	err := errors.New("no todo found")

	for key, val := range db.Todo {
		if val.ID == id {
			if val.Author == author {
				db.Todo = append(db.Todo[:key], db.Todo[key+1:]...)
				err = nil
			} else {
				err = errors.New("not authorized")
			}
		}
	}

	return err
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

func UpdateTodo(todo model.Todo, id string, author string) error {
	for key, val := range db.Todo {
		if val.ID == id {
			if val.Author != author {
				return errors.New("not authorized")
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
			return nil
		}
	}

	return errors.New("no todo found")
}

func getID() string {
	return "t10" + strconv.Itoa(len(db.Todo)+1)
}
