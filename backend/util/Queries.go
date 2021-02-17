package util

import (
	"backend/Web_project/model"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
)

func GetUserFromdbByUsername(db *sql.DB,username string)(model.User,bool,error) {
	var user model.User
	query := "SELECT * FROM user WHERE username="+"\""+username+"\""
	results, err := db.Query(query)
	if err != nil {
		return user,false,err
	}
	count:=0
	for results.Next(){
		count++
		err = results.Scan(&user.UserId,&user.UserName,&user.FirstName,&user.LastName,&user.Password)
		if err != nil {
			return user,false,err
		}
	}
	if count==0 {
		return user,false,nil
	}else {
		return user,true,nil
	}

}



func InsertNewUserIntodb(db *sql.DB,user model.User)(error){

	query := "INSERT INTO user VALUES ("+"NULL"+","+"\""+string(user.UserName)+"\""+","+"\""+string(user.FirstName)+"\""+","+"\""+string(user.LastName)+"\""+","+"\""+string(user.Password)+"\""+")"
	insert, err := db.Query(query)
	if err != nil {
		return err
	}
	defer insert.Close()
	return nil

}


func UpdateProfile(db *sql.DB,username string,firstname string,lastname string,password string)error  {

	var query string
	if password!="" {
		query = "UPDATE user SET firstname = "+"\""+firstname+"\""+" , "+"lastname = "+"\""+lastname+"\""+" , "+"password ="+"\""+password+"\"" +" WHERE username = "+"'"+username+"'"+";"
	}else {
		query = "UPDATE user SET firstname = "+"\""+firstname+"\""+" , "+"lastname = "+"\""+lastname+"\""+ " WHERE username = "+"\""+username+"\""+";"

	}

	fmt.Println(query)

	_, err := db.Query(query)
	if err != nil {

		return err
	}

	return nil
}

func SearchUser(db *sql.DB,username string)([]model.User,error)  {
	query := "SELECT username,firstname,lastname FROM user WHERE username LIKE " +"\""+username+"\""+";"
	var users []model.User
	results, err := db.Query(query)
	if err != nil {
		return nil,err
	}

	for results.Next(){
		var user model.User
		err = results.Scan(&user.UserName,&user.FirstName,&user.LastName)
		if err != nil {
			return nil,err
		}
		users=append(users,user)
	}

	return users,nil
}

func CreateNewList(db *sql.DB,username string)error  {
	query := "INSERT INTO list VALUES ("+"NULL"+","+"\""+username+"\""+")"
	insert, err := db.Query(query)
	if err != nil {
		return err
	}
	defer insert.Close()
	return nil
}

func CheckListIdExistsAndOwnerCheck(db *sql.DB,listid int,username string)bool  {
	query := "SELECT owner FROM list WHERE id="+strconv.Itoa(listid)

	results, err := db.Query(query)

	if err != nil {
		return false
	}
	count:=0
	var list model.List
	for results.Next(){
		err = results.Scan(&list.Owner)
		if err != nil {
			return false
		}
		count++
	}

	if count == 0 {
		return false
	}

	if list.Owner==username {
		return true
	}else {
		return false
	}

}

func AddItemTolist(db *sql.DB,listid int,itemid int,username string)error  {
	if CheckListIdExistsAndOwnerCheck(db,listid,username) {
		query := "INSERT INTO listitem VALUES ("+"NULL"+","+strconv.Itoa(listid)+","+strconv.Itoa(itemid)+")"
		insert, err := db.Query(query)
		if err != nil {
			return err
		}
		defer insert.Close()
		return nil
	}else {
		return errors.New("Invalid list")
	}
}




func AddFriend(db *sql.DB,sideone string,sidetwo string)error  {
	var user_sideone model.User
	var user_sidetwo model.User
	user_sideone , b1,err1:=GetUserFromdbByUsername(db,sideone)
	if err1!=nil {
		return err1
	}

	user_sidetwo , b2,err2:=GetUserFromdbByUsername(db,sidetwo)
	if err2!=nil {
		return err1
	}


	if b1&&b2 {
		querycheckduplicate := "SELECT * FROM friend WHERE sideone = "+strconv.Itoa(user_sideone.UserId) +" AND "+"sidetwo = "+strconv.Itoa(user_sidetwo.UserId)

		results, err := db.Query(querycheckduplicate)

		if err != nil {
			return err
		}
		count:=0

		for results.Next(){
			count++
			if 0<count {
				break
			}
		}
		if 0<count {
			return nil
		}
		query := "INSERT INTO friend VALUES ("+"NULL"+","+strconv.Itoa(user_sideone.UserId)+","+strconv.Itoa(user_sidetwo.UserId)+")"
		insert, err := db.Query(query)
		if err != nil {
			return err
		}
		defer insert.Close()
		return nil
	}else {
		return errors.New("User does not exists")
	}

}

func HasFriendRelation(db *sql.DB,sideone int,sidetwo int) bool {
	querycheckduplicate := "SELECT * FROM friend WHERE sideone = "+strconv.Itoa(sideone) +" AND "+"sidetwo = "+strconv.Itoa(sidetwo)
	results, err := db.Query(querycheckduplicate)
	if err != nil {
		return false
	}
	count:=0
	for results.Next(){
		count++
		if 0<count {
			break
		}
	}

	if 0<count  {
		return true
	}else {
		return false
	}
}


func BuyGift(db *sql.DB,id int,buyer string)error {

	query := "UPDATE listitem SET buystatus = 1 , buyer = \"" + buyer + "\"" + " WHERE id =" + strconv.Itoa(id)
	_, err := db.Query(query)
	if err != nil {

		return err
	}

	return nil
}

func GetList(db *sql.DB,owner string)([]model.List,error)  {
	var lists []model.List
	query := "SELECT * FROM list WHERE owner = " + "\""+ owner + "\""

	results, err := db.Query(query)
	if err != nil {
		return lists,err
	}

	for results.Next(){
		var list model.List
		err = results.Scan(&list.Id,&list.Owner)
		if err != nil {
			return lists,err
		}
		lists=append(lists,list)
	}
	return lists,nil
}


func GetListItem(db *sql.DB,listid int)([]model.ListItem,error)  {
	var lists []model.ListItem
	query := "SELECT * FROM listitem WHERE listid = " + strconv.Itoa(listid)

	results, err := db.Query(query)
	if err != nil {
		return lists,err
	}

	for results.Next(){
		var list model.ListItem
		err = results.Scan(&list.Id,&list.Listid,&list.ItemId,&list.BuyStatus,&list.Buyer)
		if err != nil {
			return lists,err
		}
		lists=append(lists,list)
	}
	return lists,nil
}

func GetItemDetail(db *sql.DB,itemid int)(model.Item,error)  {
	var item model.Item
	query := "SELECT * FROM item WHERE id = " + strconv.Itoa(itemid)
	results, err := db.Query(query)
	if err != nil {
		return item,err
	}

	for results.Next(){
		var list model.ListItem
		err = results.Scan(&list.Id,&list.Listid,&list.ItemId,&list.BuyStatus,&list.Buyer)
		if err != nil {
			return item,err
		}

	}
	return item,nil
}