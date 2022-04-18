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
	r.HandleFunc("/listallusers", Listallusers)
	r.HandleFunc("/searchanuser/{fullname}", Searchanuser).Methods(http.MethodPost, http.MethodOptions)
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

	// w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usersnamelist)
	// users, err := json.Marshal(users)
	fmt.Fprintf(w, "")
}

func Searchanuser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // 獲取引數
	user, _ := sqlpublic.Searchanuser(vars["fullname"])
	json.NewEncoder(w).Encode(user)
	fmt.Fprintf(w, "")
}
