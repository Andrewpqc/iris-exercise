package models

type Detail struct {
	Id     int    `xorm:"not null pk autoincr INTEGER"`
	Job    string `xorm:"TEXT"`
	UserId int    `xorm:"index INTEGER"`
}
