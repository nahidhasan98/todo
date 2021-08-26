package service

import (
	"errors"

	"github.com/nahidhasan98/todo/db"
	"github.com/nahidhasan98/todo/model"
)

func Authenticate(reqUser *model.Credentials) (*model.User, error) {
	for _, val := range db.Credentials {
		if reqUser.Username == val.Username && reqUser.Password == val.Password {
			return getUserInfoByUsername(val.Username), nil
		}
	}

	return nil, errors.New("wrong credentials")
}
