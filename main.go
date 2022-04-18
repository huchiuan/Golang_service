// main.go
package main

import (
	sqlpublic "UI_Assignment/sqlpublic"
	"fmt"
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func Init() {
	sqlpublic.SqlInit()
}
func main() {
	Init()
	// 連接DB

	//restfulapi
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/list-all-users", Listallusers)
	r.HandleFunc("/search-an-user/{fullname}", Searchanuser).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/get-user-detailed-information", Getuserdetailedinformation)
	http.ListenAndServe(":5000", r)

}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "hello world")
}

func Listallusers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, _ := sqlpublic.Listallusers()
	var usersnamelist []string
	for _, user := range users {
		usersnamelist = append(usersnamelist, user.Acct)
		log.Info(user.Acct)
	}

	json.NewEncoder(w).Encode(usersnamelist)
	fmt.Fprintf(w, "")
}

func Searchanuser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // 獲取引數
	user, _ := sqlpublic.Searchanuser(vars["fullname"])
	usercount := user.Acct
	json.NewEncoder(w).Encode(usercount)
	fmt.Fprintf(w, "")
}

func Getuserdetailedinformation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, _ := sqlpublic.Getuserdetailedinformation()

	json.NewEncoder(w).Encode(users)
	fmt.Fprintf(w, "")
}
