package model

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  string `json:"address"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
