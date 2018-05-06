package main

import (
	//"time"

	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	//"iris-exercise/orm/xorm/models"
	//"time"
)

func main() {
	engine, err := xorm.NewEngine("sqlite3", "./master.sqlite3")
	if err != nil {
		fmt.Println("create engine failed!")
		return
	}

	//Query返回[]map[string][]byte类型的数据
	//fmt.Println("Query返回的数据类型：")
	//results,err:=engine.Query("select * from user")
	//fmt.Println(results)
	//fmt.Println("--------------------------------------")
	//results,err=engine.Where("id=1").Query()
	//fmt.Println(results)

	//QueryString返回[]map[string]string类型的数据
	//fmt.Println("QueryString返回的数据类型：")
	//results,err:=engine.QueryString("select * from user")
	//fmt.Println(results)
	//fmt.Println("--------------------------------------")
	//results,err=engine.Where("id=1").QueryString()
	//fmt.Println(results)

	//QueryInterface返回[]map[string]interface{}类型的数据
	fmt.Println("QueryInterface返回的数据类型：")
	results, err := engine.QueryInterface("select * from user")
	fmt.Println(results)
	fmt.Println("--------------------------------------")
	results, err = engine.Where("id=1").QueryInterface()
	fmt.Println(results)

}
