package model

type Data struct {
	User User   `json:"user,omitempty"`
	Task []Todo `json:"task,omitempty"`
}

type Response struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Data    []Data `json:"data,omitempty"`
}
