package entity

import "time"

type User struct {
	Id       int
	Name     string
	Salt     string
	Age      int
	Passwd   string
	Birthday time.Time
	Created  time.Time
	Updated  time.Time
}

func (User) TableName() string {
	return "user"
}
