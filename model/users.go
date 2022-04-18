package model

type Users struct {
	Users []User `json:"Users"`
}

type User struct {
	Acct       string `gorm:"primaryKey"`
	Pwd        string `json:"pwd"`
	Fullname   string `json:"fullname"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}
