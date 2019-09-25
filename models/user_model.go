package models

import "time"

/**
models写和数据的操作 dao层
*/

type User struct {
	Id                int64
	Address           string
	Email             string
	LastActive        time.Time
	NickName          string
	Password          string
	PersonalStatement string
	PhoneNumber       string
	RegistratioDate   time.Time
	Signature         string
	Username          string
}
