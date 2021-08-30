package auth

type User struct {
	ID       string `json:"id" bson:"_id"`
	Username string `json:"username" bson:"username"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	Address  string `json:"address" bson:"address"`
}

type Token struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
