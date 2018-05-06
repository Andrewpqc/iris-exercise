//get得到单条记录

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
	queryUser := models.User{Id: 5}

	ok, err := engine.Get(&queryUser)
	if err != nil {
		fmt.Println("get user failed!")
		return
	}
	//返回是否操作成功
	fmt.Println(ok)
	fmt.Println(queryUser)

	//名字为"小1"的用户有多个，但是GET查询只会返回第一条数据
	queryUser2 := models.User{Name: "小2"}
	ok, err = engine.Get(&queryUser2)
	fmt.Println(ok)
	fmt.Println(queryUser2)

	//将查询到的数据存到user结构中
	user:=models.User{}

	_, err = engine.Get(&user)
	// SELECT * FROM user LIMIT 1

	_, err = engine.Where("name = ?", "andrew").Desc("id").Asc("name").Get(&user)
	// SELECT * FROM user WHERE name = ? ORDER BY id DESC LIMIT 1

	var name string
	_, err = engine.Where("id = ?", id).Cols("name").Get(&name)
	// SELECT name FROM user WHERE id = ?

	var id int64
	_, err = engine.Where("name = ?", name).Cols("id").Get(&id)
	_, err = engine.SQL("select id from user").Get(&id)
	// SELECT id FROM user WHERE name = ?

	var valuesMap = make(map[string]string)
	_, err = engine.Where("id = ?", id).Get(&valuesMap)
	// SELECT * FROM user WHERE id = ?

	var valuesSlice = make([]interface{}, len(cols))
	_, err = engine.Where("id = ?", id).Cols(cols...).Get(&valuesSlice)
	// SELECT col1, col2, col3 FROM user WHERE id = ?

	if err != nil {

	}


}
