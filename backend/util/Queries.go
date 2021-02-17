package util

import (
	"Web_project/model"
	"database/sql"
	"fmt"
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