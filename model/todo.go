package model

type Todo struct {
	ID      string `json:"id"`
	Task    string `json:"task"`
	At      int64  `json:"at"`
	Message string `json:"message"`
	Author  string `json:"author"`
}
