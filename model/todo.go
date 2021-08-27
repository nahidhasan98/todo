package model

type Todo struct {
	ID      string `json:"id,omitempty"`
	Task    string `json:"task,omitempty"`
	At      int64  `json:"at,omitempty"`
	Message string `json:"message,omitempty"`
	Author  string `json:"author,omitempty"`
}
