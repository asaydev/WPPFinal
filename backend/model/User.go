package model

type User struct {
	UserId int `json:"id"`
	UserName string `json:"username"`
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Password string `json:"password"`
	AcessToken string `json:"access_token"`
}