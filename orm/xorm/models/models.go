package models

import "time"


type User struct {
	Id int64 // auto-increment by-default by xorm
	Name string
	Age int
	Passwd string `xorm:"varchar(200)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

type Item struct{
	Id int64
	Name string
}
