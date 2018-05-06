package main

import (
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"iris-exercise/orm/xorm/models"
	"time"
)

func main() {
	dataSourceSlice := []string{"./master.sqlite3", "./slave1.sqlite3", "./slave2.sqlite3"}
	engineGroup, err := xorm.NewEngineGroup("sqlite3", dataSourceSlice)
	if err != nil {
		fmt.Println("create engine group failed!")
		return
	}
	err = engineGroup.Sync2(new(models.User))
	if err != nil {
		fmt.Println("sync failed!")
		return
	}

	user1 := &models.User{Name: "小1", Age: 15, Passwd: "xiaohong's password", Created: time.Now(), Updated: time.Now()}
	user2 := &models.User{Name: "小2", Age: 15, Passwd: "xiaohong's password", Created: time.Now(), Updated: time.Now()}
	user3 := &models.User{Name: "小3", Age: 15, Passwd: "xiaohong's password", Created: time.Now(), Updated: time.Now()}
	user4 := &models.User{Name: "小1", Age: 15, Passwd: "xiaohong's password", Created: time.Now(), Updated: time.Now()}

	engineGroup.Insert(user1, user2, user3, user4)
	fmt.Println("insert user1,user2,user3 to engine group successed!")

	usr1 := models.User{Name: "小1"}
	fmt.Println(engineGroup.FindAndCount(usr1))
	//engineGroup.Find(usr1)
	fmt.Println(usr1)

}
