package model

type Data struct {
	Info User   `json:"info,omitempty"`
	Task []Todo `json:"task,omitempty"`
}

type Response struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Data    []Data `json:"data,omitempty"`
}
