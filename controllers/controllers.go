package controllers

import (
	Auth "UI_Assignment/auth"
	Users "UI_Assignment/model"
	sqlpublic "UI_Assignment/sqlpublic"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
)

//index
var IndexHandler = func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "hello world")
}

//commit 1: Create an API to list all users.
var Listallusers = func(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	users, _ := sqlpublic.Listallusers()
	var usersnamelist []string
	for _, user := range users {
		usersnamelist = append(usersnamelist, user.Acct)
		log.Info(user.Acct)
	}

	json.NewEncoder(w).Encode(usersnamelist)
	fmt.Fprintf(w, "sussess_token:")
}

//commit 2: Create an API to search an user by fullname.
var Searchanuser = func(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	log.Println(dec)
	dec.DisallowUnknownFields()
	var user = Users.User{}
	err := dec.Decode(&user)
	if err != nil {
		panic(err)
	}

	user, _ = sqlpublic.Searchanuser(user.Fullname)
	usercount := user.Acct
	json.NewEncoder(w).Encode(usercount)
	fmt.Fprintf(w, "")
}

//commit 3: Create an API to get the user’s detailed information.
var Getuserdetailedinformation = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, _ := sqlpublic.Getuserdetailedinformation()

	json.NewEncoder(w).Encode(users)
	fmt.Fprintf(w, "")
}

//commit 4: Create an API to create the user (user sign up).
var Createtheuser = func(w http.ResponseWriter, r *http.Request) {

	dec := json.NewDecoder(r.Body)
	log.Println(dec)

	dec.DisallowUnknownFields()
	var user = Users.User{
		Created_at: time.Now(),
		Updated_at: time.Now(),
	}
	err := dec.Decode(&user)

	if err != nil {
		panic(err)
	}
	log.Println(user)

	_ = sqlpublic.Createtheuser(user)

	usercount := user
	json.NewEncoder(w).Encode(usercount)
	fmt.Fprintf(w, "")
}

// type Claims struct {
// 	Acct string `json:"acct"`
// 	jwt.StandardClaims
// }

//commit 5: Create an API to generate the token to the user (user sign in).
var Signin = func(w http.ResponseWriter, r *http.Request) {

	dec := json.NewDecoder(r.Body)
	log.Println(dec)
	dec.DisallowUnknownFields()
	var user = Users.User{}
	err := dec.Decode(&user)
	if err != nil {
		panic(err)
	}

	checkuser, _ := sqlpublic.Getuserbyacct(user.Acct)

	realPwd := checkuser.Pwd
	inputPwd := user.Pwd

	if realPwd == inputPwd {

		//jwt
		expirationTime := time.Now().Add(5 * time.Minute)

		claims := &Auth.Claims{
			Acct: checkuser.Acct,
			StandardClaims: jwt.StandardClaims{
				// In JWT, the expiry time is expressed as unix milliseconds
				ExpiresAt: expirationTime.Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(Auth.JwtKey)
		if err != nil {
			// If there is an error in creating the JWT return an internal server error
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "密碼正確")

	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "密碼錯誤")
	}
}

//commit 6: Create an API to delete the user.
var Deletetheuser = func(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	log.Println(dec)
	dec.DisallowUnknownFields()
	var user = Users.User{}
	err := dec.Decode(&user)
	if err != nil {
		panic(err)
	}

	_ = sqlpublic.Deleteuserbyacct(user.Acct)
	// //
	// usercount := user.Acct
	// json.NewEncoder(w).Encode(usercount)
	fmt.Fprintf(w, "刪除結束")
}

//commit 7: Create an API to update the user.
var Updatetheuser = func(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	log.Println(dec)
	dec.DisallowUnknownFields()
	var user = Users.User{}
	err := dec.Decode(&user)
	if err != nil {
		panic(err)
	}

	_ = sqlpublic.Updateuser(user.Acct, user.Pwd)
	// //
	// usercount := user.Acct
	// json.NewEncoder(w).Encode(usercount)
	fmt.Fprintf(w, "更新結束")
}

//commit 9: Create an API to update user’s fullname.
var Updatetheuserfullname = func(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	log.Println(dec)
	dec.DisallowUnknownFields()
	var user = Users.User{}
	err := dec.Decode(&user)
	if err != nil {
		panic(err)
	}

	_ = sqlpublic.Updateuserfullname(user.Acct, user.Fullname)
	// //
	// usercount := user.Acct
	// json.NewEncoder(w).Encode(usercount)
	fmt.Fprintf(w, "更新結束")
}
