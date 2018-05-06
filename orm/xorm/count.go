package main

import (
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"iris-exercise/orm/xorm/models"
)

func main() {
	engine, err := xorm.NewEngine("sqlite3", "./master.sqlite3")
	if err != nil {
		fmt.Println("create engine failed!")
		return
	}

	counts, _ := engine.Count(&models.User{Id: 1})
	fmt.Println(counts)
	counts2, _ := engine.Count(&models.User{Name: "小1"})
	fmt.Println(counts2)

}
