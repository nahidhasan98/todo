package user

type User struct {
	ID       string `json:"id" bson:"_id"`
	Username string `json:"username" bson:"username"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	Address  string `json:"address" bson:"address"`
}

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

type Data struct {
	User User   `json:"user"`
	Task []Todo `json:"task"`
}
type UserResponse struct {
	UData   *[]Data `json:"data"`
	Err     string  `json:"err"`
	Message string  `json:"message"`
}
