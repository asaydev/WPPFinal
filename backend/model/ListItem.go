package model

type ListItem struct {
	Id int `json:"id"`
	Listid  int `json:"listid"`
	ItemId  int `json:"itemid"`
	BuyStatus  int `json:"buystatus"`
	Buyer string `json:"buyer"`
}



