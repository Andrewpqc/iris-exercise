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

	//从user表中删除id为1的数据,这是一次删除一条数据
	engine.ID(1).Delete(&models.User{})
	//从user表中删除name为updateName的数据，这里是一次删除多条数据
	engine.Where("name=?", "updateName").Delete(&models.User{})

}
