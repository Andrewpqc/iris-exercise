package main

import (
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"iris-exercise/orm/xorm/models"
	"time"
)

func main() {
	engine, err := xorm.NewEngine("sqlite3", "./master.sqlite3")
	if err != nil {
		fmt.Println("create engine failed!")
		return
	}
	user1 := &models.User{Name: "andrew", Age: 22, Created: time.Now(), Updated: time.Now()}
	user2 := &models.User{Name: "andrew2", Age: 22, Created: time.Now(), Updated: time.Now()}
	user3 := &models.User{Name: "andrew3", Age: 22, Created: time.Now(), Updated: time.Now()}
	user4 := &models.User{Name: "andrew4", Age: 22, Created: time.Now(), Updated: time.Now()}

	//这里的affected是返回被成功插入的数据
	//插入一条数据
	affected, err := engine.Insert(user1)
	if err != nil {
		fmt.Println("%v", err)
		return
	}
	fmt.Println(affected)
	//一次插入多条数据
	affected, err = engine.Insert(user2, user3, user4)
	if err != nil {
		fmt.Println("%v", err)
		return
	}
	fmt.Println(affected)
	fmt.Println("inser successed!")

}
