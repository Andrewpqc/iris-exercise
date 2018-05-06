package main

import (
	//"time"

	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"iris-exercise/orm/xorm/models"

	"time"
)

func main() {
	var driverName = "sqlite3"
	var dataSourceName = "./master.sqlite3"
	engine, err := xorm.NewEngine(driverName, dataSourceName)
	if err != nil {
		fmt.Println("create engine failed!")
		return
	}

	err = engine.Sync2(new(models.User))
	if err != nil {
		fmt.Println("create table failed!")
		return
	}

	fmt.Println("create engine and sync data successed!")
	user := &models.User{Name: "小明", Passwd: "xiaoming's password", Age: 13, Created: time.Now(), Updated: time.Now()}
	engine.Insert(user)

	fmt.Println("create a entity of user 小明")

	queryuser := models.User{Id: 1}
	if ok, _ := engine.Get(&queryuser); ok {
		fmt.Println("find the user that id =1:")
		fmt.Println(queryuser)
	}
}
