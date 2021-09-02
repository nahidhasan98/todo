package todo

type Todo struct {
	ID      string `json:"id" bson:"_id"`
	Task    string `json:"task" bson:"task"`
	At      int64  `json:"at" bson:"at"`
	Message string `json:"message" bson:"message"`
	Author  string `json:"author" bson:"author"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
