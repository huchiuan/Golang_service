package model

import "time"

type Users struct {
	Users []User `json:"Users"`
}

type User struct {
	Acct       string    `gorm:"primaryKey"`
	Pwd        string    `json:"pwd"`
	Fullname   string    `json:"fullname"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
