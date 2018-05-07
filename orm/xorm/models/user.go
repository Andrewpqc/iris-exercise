package models

import (
	"time"
)

type User struct {
	Id      int       `xorm:"not null pk autoincr INTEGER"`
	Name    string    `xorm:"TEXT"`
	Age     int       `xorm:"INTEGER"`
	Passwd  string    `xorm:"TEXT"`
	Created time.Time `xorm:"DATETIME"`
	Updated time.Time `xorm:"DATETIME"`
}
