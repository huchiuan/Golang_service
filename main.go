// main.go
package main

import (
	"UI_Assignment/auth"
	controllers "UI_Assignment/controllers"
	sqlpublic "UI_Assignment/sqlpublic"
	"net/http"

	// "encoding/json"

	// Users "UI_Assignment/model"
	// "time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	// log "github.com/sirupsen/logrus"
)

func Init() {
	sqlpublic.SqlInit()
}

func main() {
	Init()
	// 連接DB

	//restfulapi
	r := mux.NewRouter()

	// 註冊jwt認證的中介軟體
	r.Use(auth.JwtAuthentication)
	r.HandleFunc("/", controllers.IndexHandler)
	r.HandleFunc("/list-all-users", controllers.Listallusers)
	r.HandleFunc("/search-an-user", controllers.Searchanuser).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/get-user-detailed-information", controllers.Getuserdetailedinformation)
	r.HandleFunc("/create-the-user", controllers.Createtheuser).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/sign-in", controllers.Signin).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/delete-the-user", controllers.Deletetheuser).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/update-the-user", controllers.Updatetheuser).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/update-the-user-fullname", controllers.Updatetheuserfullname).Methods(http.MethodPost, http.MethodOptions)

	http.ListenAndServe(":5000", r)

}
