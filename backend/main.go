package main

import (
	"backend/Web_project/model"
	"backend/Web_project/util"
	"encoding/json"
	"fmt"
	"github.com/gobuffalo/uuid"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

var db = util.Initializedatabase()
var port = "8080"
var dbSessions = map[string]model.Session{} // session ID, session
const sessionLength int = 5*60*60

func main() {
	defer db.Close()
	fmt.Println("Server Starts Successfully !")
	handleRequests()
}


func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/register", signup)
	router.HandleFunc("/login", signin)
	//router.HandleFunc("/updateprofile", authorized(updateprofile))
	router.HandleFunc("/updateprofile", updateprofile)
	//router.HandleFunc("/api/searchuser", authorized(searchuser))
	router.HandleFunc("/searchuser", searchuser)
	router.HandleFunc("/api/createlist", authorized(createlist))
	//router.HandleFunc("/api/additemtolist", authorized(additemtolist))
	router.HandleFunc("/additemtolist", additemtolist)
	//router.HandleFunc("/api/addfriend", authorized(addfriend))
	router.HandleFunc("/addfriend", addfriend)
	//router.HandleFunc("/api/buygift", authorized(buygift))
	router.HandleFunc("/buygift", buygift)
	//router.HandleFunc("/api/getlist", authorized(getList))
	router.HandleFunc("/getlist", getList)
	//router.HandleFunc("/api/getlistitem", authorized(getListItem))
	router.HandleFunc("/getlistitem", getListItem)
	//router.HandleFunc("/api/getitemdetail", authorized(getListItem))
	router.HandleFunc("/getitemdetail", getitemdetail)

	http.ListenAndServe(":"+port, router)
}

func alreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	cookie, err := req.Cookie("session")
	if err != nil {
		return false
	}
	session, ok := dbSessions[cookie.Value]
	if ok {
		session.LastActivity = time.Now()
		dbSessions[cookie.Value] = session
	}

	user,ok,err:=util.GetUserFromdbByUsername(db,session.Un)
	if err!=nil {
		fmt.Println(err)
		return false
	}
	// refresh session
	cookie.MaxAge = sessionLength
	http.SetCookie(w, cookie)
	return ok&&("" != user.UserName)
}


func authorized(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !alreadyLoggedIn(w, r) {
			http.Error(w, "not logged in", http.StatusUnauthorized)
			return // don't call original handler
		}
		h.ServeHTTP(w, r)
	})
}

func getHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
func signup(w http.ResponseWriter, req *http.Request) {


	if alreadyLoggedIn(w, req) {
		w.WriteHeader(http.StatusSeeOther)
		return
	}

	// process form submission
	if req.Method == http.MethodPost {
		//get json values
		out := make([]byte,1024)

		bodyLen, err := req.Body.Read(out)

		if err != io.EOF {
			fmt.Println(err.Error())
			w.Write([]byte("{error:" + err.Error() + "}"))
			return
		}
		var k model.User

		err = json.Unmarshal(out[:bodyLen],&k)

		if err != nil {
			w.Write([]byte("{error:" + err.Error() + "}"))
			return
		}
		fmt.Println(k.UserName)
		fmt.Println(k.FirstName)
		fmt.Println(k.LastName)
		fmt.Println(k.Password)
		//null value
		if k.UserName=="" || k.FirstName=="" || k.LastName=="" || k.Password=="" {
			http.Error(w, "Please fill all the variables", http.StatusSeeOther)
			return
		}
		// username taken?
		user,ok,err:=util.GetUserFromdbByUsername(db,k.UserName)
		if ok&&(user.UserName!="") {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		if err!=nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// create session
		sID, _ := uuid.NewV4()
		cookie := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		cookie.MaxAge = sessionLength
		http.SetCookie(w, cookie)
		dbSessions[cookie.Value] = model.Session{k.UserName, time.Now()}
		// store user in dbUsers
		bs, err := bcrypt.GenerateFromPassword([]byte(k.Password), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		user = model.User{
			UserName:k.UserName,
			Password:string(bs),
			FirstName:k.FirstName,
			LastName:k.LastName,
		}

		err = util.InsertNewUserIntodb(db,user)
		if err!=nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)

	}else {
		w.WriteHeader(http.StatusBadRequest)
	}
}


func signin(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(w, req) {
		w.WriteHeader(http.StatusSeeOther)
		return
	}
	// process form submission
	if req.Method == http.MethodPost {
		out := make([]byte,1024)

		bodyLen, err := req.Body.Read(out)

		if err != io.EOF {
			fmt.Println(err.Error())
			w.Write([]byte("{error:" + err.Error() + "}"))
			return
		}
		var k model.User

		err = json.Unmarshal(out[:bodyLen],&k)

		if err != nil {
			w.Write([]byte("{error:" + err.Error() + "}"))
			return
		}

		fmt.Println(k.UserName)
		fmt.Println(k.Password)
		// is there a username?
		user,ok,err := util.GetUserFromdbByUsername(db,k.UserName)

		if err!=nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		if !ok&&user.UserName=="" {
			//w.WriteHeader(http.StatusOK)
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// does the entered password match the stored password?

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(k.Password))
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// create session
		sID, _ := uuid.NewV4()
		cookie := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		cookie.MaxAge = sessionLength
		http.SetCookie(w, cookie)
		dbSessions[cookie.Value] = model.Session{k.UserName, time.Now()}
		w.WriteHeader(http.StatusOK)

	}else {
		w.WriteHeader(http.StatusBadRequest)
	}

}


func updateprofile(w http.ResponseWriter, req *http.Request) {

	// process form submission
	if req.Method == http.MethodPut {
		out := make([]byte,1024)

		bodyLen, err := req.Body.Read(out)

		if err != io.EOF {
			fmt.Println(err.Error())
			w.Write([]byte("{error:" + err.Error() + "}"))
			return
		}
		var k model.UserUpdate

		err = json.Unmarshal(out[:bodyLen],&k)

		if err != nil {
			w.Write([]byte("{error:" + err.Error() + "}"))
			return
		}


		cookie, err := req.Cookie("session")
		if err != nil {
			w.WriteHeader(http.StatusSeeOther)
			return
		}
		session, ok := dbSessions[cookie.Value]
		if ok {
			session.LastActivity = time.Now()
			dbSessions[cookie.Value] = session
		}
		var repass bool
		if k.Password!="" && k.RePassword!="" {
			user,_,err:=util.GetUserFromdbByUsername(db,session.Un)
			err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(k.Password))
			if err != nil {
				http.Error(w, "Password is not correct", http.StatusForbidden)
				repass = false
				return
			}else {
				repass = true
			}
		}

		if repass {
			bs, err := bcrypt.GenerateFromPassword([]byte(k.Password), bcrypt.MinCost)
			if err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}
			err = util.UpdateProfile(db,session.Un,k.FirstName,k.LastName,string(bs))
			if err != nil {

				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}else {
			err := util.UpdateProfile(db,session.Un,k.FirstName,k.LastName,"")
			if err != nil {

				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		w.WriteHeader(http.StatusOK)

	}else {
		w.WriteHeader(http.StatusBadRequest)
	}

}

func searchuser(w http.ResponseWriter, req *http.Request)  {

	if req.Method == http.MethodGet {
		out := make([]byte,1024)

		bodyLen, err := req.Body.Read(out)

		if err != io.EOF {
			fmt.Println(err.Error())
			w.Write([]byte("{error:" + err.Error() + "}"))
			return
		}
		var k model.User

		err = json.Unmarshal(out[:bodyLen],&k)

		if err != nil {
			w.Write([]byte("{error:" + err.Error() + "}"))
			return
		}

		results,err := util.SearchUser(db,k.UserName)
		if err!=nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(results)
	}else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func createlist(w http.ResponseWriter, req *http.Request)  {
	cookie, err := req.Cookie("session")
	if err != nil {
		w.WriteHeader(http.StatusSeeOther)
		return
	}

	session, ok := dbSessions[cookie.Value]
	if ok {
		session.LastActivity = time.Now()
		dbSessions[cookie.Value] = session
	}

	err = util.CreateNewList(db,session.Un)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func additemtolist(w http.ResponseWriter, req *http.Request)  {
	if req.Method == http.MethodPost {
		out := make([]byte,1024)

		bodyLen, err := req.Body.Read(out)

		if err != io.EOF {
			fmt.Println(err.Error())
			w.Write([]byte("{error:" + err.Error() + "}"))
			return
		}
		var k model.ListItem

		err = json.Unmarshal(out[:bodyLen],&k)

		if err != nil {
			w.Write([]byte("{error:" + err.Error() + "}"))
			return
		}
		cookie, err := req.Cookie("session")
		if err != nil {
			w.WriteHeader(http.StatusSeeOther)
			return
		}

		session, ok := dbSessions[cookie.Value]
		if ok {
			session.LastActivity = time.Now()
			dbSessions[cookie.Value] = session
		}

		listid_int, err := strconv.Atoi(string(k.Listid))
		if err != nil {
			http.Error(w, "listid value is not valid", http.StatusBadRequest)
			return
		}

		itemid_int, err := strconv.Atoi(string(k.ItemId))
		if err != nil {
			http.Error(w, "itemid value is not valid", http.StatusBadRequest)
			return
		}

		err = util.AddItemTolist(db,listid_int,itemid_int,session.Un)

		if err!=nil {
			http.Error(w, err.Error(), http.StatusSeeOther)
			return
		}

		w.WriteHeader(http.StatusOK)
	}else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func addfriend(w http.ResponseWriter, req *http.Request)  {
	out := make([]byte,1024)

	bodyLen, err := req.Body.Read(out)

	if err != io.EOF {
		fmt.Println(err.Error())
		w.Write([]byte("{error:" + err.Error() + "}"))
		return
	}
	var k model.UserFriend


	err = json.Unmarshal(out[:bodyLen],&k)

	if err != nil {
		w.Write([]byte("{error:" + err.Error() + "}"))
		return
	}
	cookie, err := req.Cookie("session")
	if err != nil {
		w.WriteHeader(http.StatusSeeOther)
		return
	}

	session, ok := dbSessions[cookie.Value]
	if ok {
		session.LastActivity = time.Now()
		dbSessions[cookie.Value] = session
	}

	if k.Friend==session.Un {
		http.Error(w, err.Error(), http.StatusSeeOther)
		return
	}
	err = util.AddFriend(db,session.Un,k.Friend)

	if err!=nil {
		http.Error(w, err.Error(), http.StatusSeeOther)
		return
	}

	w.WriteHeader(http.StatusOK)
}


func buygift(w http.ResponseWriter, req *http.Request)  {
	out := make([]byte,1024)

	bodyLen, err := req.Body.Read(out)

	if err != io.EOF {
		fmt.Println(err.Error())
		w.Write([]byte("{error:" + err.Error() + "}"))
		return
	}
	var k model.Item


	err = json.Unmarshal(out[:bodyLen],&k)

	if err != nil {
		w.Write([]byte("{error:" + err.Error() + "}"))
		return
	}

	cookie, err := req.Cookie("session")
	if err != nil {
		w.WriteHeader(http.StatusSeeOther)
		return
	}

	session, ok := dbSessions[cookie.Value]
	if ok {
		session.LastActivity = time.Now()
		dbSessions[cookie.Value] = session
	}


	id_int, err := strconv.Atoi(string(k.Id))
	if err != nil {
		http.Error(w, "id value is not valid", http.StatusBadRequest)
		return
	}

	err=util.BuyGift(db,id_int,session.Un)

	if err!=nil {
		http.Error(w, err.Error(), http.StatusSeeOther)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func getList(w http.ResponseWriter, req *http.Request)  {

	if req.Method == http.MethodGet{
		out := make([]byte,1024)

		bodyLen, err := req.Body.Read(out)

		if err != io.EOF {
			fmt.Println(err.Error())
			w.Write([]byte("{error:" + err.Error() + "}"))
			return
		}
		var k model.User


		err = json.Unmarshal(out[:bodyLen],&k)

		if err != nil {
			w.Write([]byte("{error:" + err.Error() + "}"))
			return
		}
		cookie, err := req.Cookie("session")
		if err != nil {
			w.WriteHeader(http.StatusSeeOther)
			return
		}

		session, ok := dbSessions[cookie.Value]
		if ok {
			session.LastActivity = time.Now()
			dbSessions[cookie.Value] = session
		}
		lists,err := util.GetList(db,k.UserName)
		json.NewEncoder(w).Encode(lists)
	}else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func getListItem(w http.ResponseWriter, req *http.Request)  {

	if req.Method == http.MethodGet{
		out := make([]byte,1024)

		bodyLen, err := req.Body.Read(out)

		if err != io.EOF {
			fmt.Println(err.Error())
			w.Write([]byte("{error:" + err.Error() + "}"))
			return
		}
		var k model.ListItem


		err = json.Unmarshal(out[:bodyLen],&k)

		if err != nil {
			w.Write([]byte("{error:" + err.Error() + "}"))
			return
		}


		listid_int, err := strconv.Atoi(string(k.Listid))

		if err != nil {
			http.Error(w, "id value is not valid", http.StatusBadRequest)
			return
		}

		cookie, err := req.Cookie("session")
		if err != nil {
			w.WriteHeader(http.StatusSeeOther)
			return
		}

		session, ok := dbSessions[cookie.Value]
		if ok {
			session.LastActivity = time.Now()
			dbSessions[cookie.Value] = session
		}
		lists,err := util.GetListItem(db,listid_int)
		if err!=nil {
			http.Error(w, err.Error(), http.StatusSeeOther)
			return
		}
		json.NewEncoder(w).Encode(lists)
	}else {
		w.WriteHeader(http.StatusBadRequest)
	}

}

func getitemdetail(w http.ResponseWriter, req *http.Request)  {

	if req.Method == http.MethodGet{
		out := make([]byte,1024)

		bodyLen, err := req.Body.Read(out)

		if err != io.EOF {
			fmt.Println(err.Error())
			w.Write([]byte("{error:" + err.Error() + "}"))
			return
		}
		var k model.Item


		err = json.Unmarshal(out[:bodyLen],&k)
			
		if err != nil {
			w.Write([]byte("{error:" + err.Error() + "}"))
			return
		}


		itemid_int, err := strconv.Atoi(string(k.Id))
		if err != nil {

			http.Error(w, "id value is not valid", http.StatusBadRequest)
			return
		}

		item,err:=util.GetItemDetail(db,itemid_int)
		if err!=nil {
			http.Error(w, err.Error(), http.StatusSeeOther)
			return
		}
		json.NewEncoder(w).Encode(item)
	}
}
