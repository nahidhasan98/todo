package service

import (
	"errors"

	"github.com/nahidhasan98/todo/db"
	"github.com/nahidhasan98/todo/model"
)

func GetAllUser(err error, isAuth bool) model.Response {
	data := model.Response{
		Status: "success",
	}
	if !isAuth {
		data.Message = err.Error() + ". limited access"
	}

	for _, val := range db.Users {
		temp := model.Data{}
		temp.User = val

		if isAuth {
			temp.Task = getTaskByUsername(val.Username)
		}

		data.Data = append(data.Data, temp)
	}

	return data
}

func getTaskByUsername(author string) []model.Todo {
	data := []model.Todo{}
	for _, val := range db.Todo {
		if val.Author == author {
			data = append(data, val)
		}
	}

	return data
}

func GetSingleUser(id string, err error, isAuth bool) (model.Response, error) {
	data := model.Response{
		Status: "success",
	}
	if !isAuth {
		data.Message = err.Error() + ". limited access"
	}

	for _, val := range db.Users {
		if val.ID == id {
			temp := model.Data{}
			temp.User = val

			if isAuth {
				temp.Task = getTaskByUsername(val.Username)
			}

			data.Data = append(data.Data, temp)
			return data, nil
		}
	}

	return data, errors.New("no user found")
}

func getUserInfoByUsername(username string) *model.User {
	for _, val := range db.Users {
		if username == val.Username {
			return &val
		}
	}
	return nil
}
