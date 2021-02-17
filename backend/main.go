package main

import (
	"Web_project/model"
	"Web_project/util"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)
var jwtKey = []byte("my_secret_key")
var signinTokens []string
var signupTokens [] string
var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

//
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
var db = util.Initializedatabase()
var port = "8080"
var dbSessions = map[string]model.Session{} // session ID, session
const sessionLength int = 5*60*60

func main() {
	//defer db.Close()
	fmt.Println("Server Starts Successfully !")
	handleRequests()
}


func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/signup", signup)
	router.HandleFunc("/api/signin", signin)
	router.HandleFunc("/api/updateprofile", authorized(updateprofile))
	router.HandleFunc("/api/searchuser", authorized(searchuser))
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


func signup(w http.ResponseWriter, req *http.Request) {

	var creds Credentials

	expiration_time := time.Now().Add(60 * time.Minute)
	claims := &Claims{
		Username: creds.Username,
		StandardClaims : jwt.StandardClaims{
			ExpiresAt: expiration_time.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	fmt.Printf(tokenString)


	signupTokens = append(signupTokens, tokenString)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if alreadyLoggedIn(w, req) {
		w.WriteHeader(http.StatusSeeOther)
		return
	}

	// process form submission
	if req.Method == http.MethodPost {
		// get form values
		username := req.FormValue("username")
		firstname := req.FormValue("firstname")
		lastname := req.FormValue("lastname")
		password := req.FormValue("password")


		//null value
		if username=="" || firstname=="" || lastname=="" || password=="" {
			http.Error(w, "Please fill all the variables", http.StatusSeeOther)
			return
		}
		// username taken?
		user,ok,err:=util.GetUserFromdbByUsername(db,username)
		if ok&&(user.UserName!="") {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		if err!=nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// create session
		//sID, _ := uuid.NewV4()
		//cookie := &http.Cookie{
		//	Name:  "session",
		//	Value: sID.String(),
		//}
		//cookie.MaxAge = sessionLength
		//http.SetCookie(w, cookie)
		//dbSessions[cookie.Value] = model.Session{username, time.Now()}
		// store user in dbUsers
		bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		user = model.User{
			UserName:username,
			Password:string(bs),
			FirstName:firstname,
			LastName:lastname,
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
	var creds Credentials
	//creat token with 1 hour expiration time
	expiration_time := time.Now().Add(60 * time.Minute)
	claims := &Claims{
		Username: creds.Username,
		StandardClaims : jwt.StandardClaims{
			ExpiresAt: expiration_time.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	fmt.Printf(tokenString)


	signinTokens = append(signinTokens, tokenString)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if alreadyLoggedIn(w, req) {
		w.WriteHeader(http.StatusSeeOther)
		return
	}
	// process form submission
	if req.Method == http.MethodPost {
		username := req.FormValue("username")
		password := req.FormValue("password")
		// is there a username?
		user,ok,err := util.GetUserFromdbByUsername(db,username)

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

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}

		if err != nil {
			// If there is an error in creating the JWT return an internal server error
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		//http.SetCookie(w, &http.Cookie{
		//	Name:    "token",
		//	Value:   tokenString,
		//	Expires: expiration_time,
		//})

			// create session
		//sID, _ := uuid.NewV4()
		//cookie := &http.Cookie{
		//	Name:  "session",
		//	Value: sID.String(),
		//}
		//cookie.MaxAge = sess	ionLength
		//http.SetCookie(w, cookie)
		//dbSessions[cookie.Value] = model.Session{username, time.Now()}
		//w.WriteHeader(http.StatusOK)
	}else {
		w.WriteHeader(http.StatusBadRequest)
	}

}


func updateprofile(w http.ResponseWriter, req *http.Request) {

	// process form submission
	if req.Method == http.MethodPut {
		firstname := req.FormValue("firstname")
		lastname := req.FormValue("lastname")
		password := req.FormValue("password")
		repassword := req.FormValue("repassword")
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
		if password!="" && repassword!="" {
			user,_,err:=util.GetUserFromdbByUsername(db,session.Un)
			err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
			if err != nil {
				http.Error(w, "Password is not correct", http.StatusForbidden)
				repass = false
				return
			}else {
				repass = true
			}
		}

		if repass {
			bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
			if err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}
			err = util.UpdateProfile(db,session.Un,firstname,lastname,string(bs))
			if err != nil {

				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}else {
			err := util.UpdateProfile(db,session.Un,firstname,lastname,"")
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
		username := req.FormValue("username")
		results,err := util.SearchUser(db,username)
		if err!=nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(results)
	}else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

