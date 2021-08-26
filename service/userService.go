package service

import (
	"errors"

	"github.com/nahidhasan98/todo/db"
	"github.com/nahidhasan98/todo/model"
)

func GetAllUser() []model.User {
	return db.Users
}

func GetSingleUser(id string) (*model.User, error) {
	for _, val := range db.Users {
		if id == val.ID {
			return &val, nil
		}
	}
	return nil, errors.New("no user found")
}

func getUserInfoByUsername(username string) *model.User {
	for _, val := range db.Users {
		if username == val.Username {
			return &val
		}
	}
	return nil
}
