//find用于查找多条数据

package main

import (
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"iris-exercise/orm/xorm/models"

	"fmt"
)

func main() {
	engine, err := xorm.NewEngine("sqlite3", "./master.sqlite3")
	if err != nil {
		fmt.Println("create engine failed!")
		return
	}

	var users []models.User
	engine.Where("name=?", "小1").And("age>0").Find(&users)

	for _, user := range users {
		fmt.Println(user)
	}

}
