package main

import (
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	engine, err := xorm.NewEngine("sqlite3", "./master.sqlite3")
	if err != nil {
		fmt.Println("create engine failed!")
		return
	}

	affected, err := engine.Exec("update user set age = ? where name = ?", 1000, "小1")
	if err != nil {
		fmt.Println("exec the sql failed!")
		return
	}

	fmt.Println("the affected record:")
	fmt.Println(affected.LastInsertId())
	fmt.Println(affected.RowsAffected()) //返回被影响的数据库列数

}
