package models

type Item struct {
	Id       int     `xorm:"not null pk autoincr INTEGER"`
	ItemName string  `xorm:"TEXT"`
	Price    float32 `xorm:"REAL"`
	Num      int     `xorm:"INTEGER"`
}
