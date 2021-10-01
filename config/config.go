package config

import "os"

var DbConnectionString = os.Getenv("MongoDB_host")

const (
	DBName    = "todo"
	AuthTable = "user"
	TodoTable = "todo"
	UserTable = "user"
)
