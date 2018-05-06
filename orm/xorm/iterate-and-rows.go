//iterate,rows都是对元组进行迭代的函数

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

	//对每一个符合条件的元组调用回调函数
	engine.Iterate(&models.User{Name: "andrew"}, func(idx int, bean interface{}) error {
		user := bean.(*models.User) //这里得到的是每一条选出来的元组
		//可以对这个元组做一些自己想做的修改，这里仅仅打印出元组
		fmt.Println(user)
		return nil
	})

	fmt.Println("------------------------------")

	engine.BufferSize(3).Iterate(&models.User{Name: "小1"}, func(idx int, bean interface{}) error {
		user := bean.(*models.User)
		fmt.Println(user)
		return nil
	})

	fmt.Println("result from rows")
	rows, err := engine.Rows(&models.User{Name: "小1"})
	// SELECT * FROM user
	defer rows.Close()
	bean := new(models.User)
	for rows.Next() {
		err = rows.Scan(bean)
		fmt.Println(bean)
	}
}
