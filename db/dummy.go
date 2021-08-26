package db

import "github.com/nahidhasan98/todo/model"

var Users = []model.User{
	{
		ID:       "u101",
		Username: "nahid",
		Email:    "nahid@mail.com",
		Address:  "Dhaka",
	},
	{
		ID:       "u102",
		Username: "hasan",
		Email:    "hasan@mail.com",
		Address:  "Kushtia",
	},
	{
		ID:       "u103",
		Username: "alex",
		Email:    "alex@mail.com",
		Address:  "Manchester",
	},
}

var Credentials = []model.Credentials{
	{
		Username: "nahid",
		Password: "123",
	},
	{
		Username: "hasan",
		Password: "456",
	},
	{
		Username: "alex",
		Password: "789",
	},
}

var Todo = []model.Todo{
	{
		ID:      "t101",
		Task:    "task 1",
		At:      1629955312,
		Message: "very important",
		Author:  "nahid",
	},
	{
		ID:      "t102",
		Task:    "task 2",
		At:      1629957512,
		Message: "",
		Author:  "nahid",
	},
	{
		ID:      "t103",
		Task:    "task 1",
		At:      1629958512,
		Message: "important",
		Author:  "hasan",
	},
}
