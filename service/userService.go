package service

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nahidhasan98/todo/db"
	"github.com/nahidhasan98/todo/model"
)

func getTaskByUsername(author string) []model.Todo {
	data := []model.Todo{}
	for _, val := range db.Todo {
		if val.Author == author {
			data = append(data, val)
		}
	}

	return data
}

func GetAllUserData(isAuth bool) ([]model.Data, error) {
	data := []model.Data{}

	if len(db.Users) == 0 {
		return nil, errors.New("user list empty")
	}

	for _, val := range db.Users {
		temp := model.Data{}
		temp.User = val

		if isAuth {
			temp.Task = getTaskByUsername(val.Username)
		}

		data = append(data, temp)
	}

	return data, nil
}

func GetSingleUserData(id string, isAuth bool) ([]model.Data, int, error) {
	var data []model.Data

	if len(db.Users) == 0 {
		return nil, http.StatusOK, errors.New("user list empty")
	}

	for _, val := range db.Users {
		if val.ID == id {
			temp := model.Data{}
			temp.User = val

			if isAuth {
				temp.Task = getTaskByUsername(val.Username)
			}

			data = append(data, temp)

			return data, http.StatusOK, nil
		}
	}

	return data, http.StatusBadRequest, errors.New("no user found")
}

func GetUserData(ctx *gin.Context, err error) (int, model.Response) {
	data := model.Response{
		Status: "success",
	}
	userID := ctx.Param("id")
	var statusCode int
	var temp []model.Data

	if err != nil {
		if userID == "" {
			temp, _ = GetAllUserData(false)
		} else {
			temp, _, _ = GetSingleUserData(userID, false)
		}

		data.Message = err.Error() + ". limited access"
		statusCode = http.StatusUnauthorized
	} else {
		var err2 error
		if userID == "" {
			temp, err2 = GetAllUserData(true)
		} else {
			temp, statusCode, err2 = GetSingleUserData(userID, true)
		}

		if err2 != nil {
			data.Message = err2.Error()
		}
	}

	data.Data = temp

	return statusCode, data
}

func getUserInfoByUsername(username string) *model.User {
	for _, val := range db.Users {
		if username == val.Username {
			return &val
		}
	}

	return nil
}
