package model

type User struct {
	UserId int `json:"id"`
	UserName string `json:"username"`
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Password string `json:"password"`
}

type UserUpdate struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Password string `json:"password"`
	RePassword string `json:"repassword"`
}
type UserFriend struct {
	Friend string `json:"friend"`
}


