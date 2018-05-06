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

	//判断该数据库中是否存在这张表
	has, err := engine.Exist(new(models.User))
	fmt.Println(has)

	has, err = engine.Exist(new(models.Item))
	fmt.Println(has)

	//判断数据库中是否存在一个名为andrew的人
	has, err = engine.Exist(&models.User{
		Name: "andrew",
	})
	fmt.Println(has)

	has, err = engine.Exist(&models.User{
		Name: "小道上是否",
		Age:  10000,
	})

	fmt.Println(has)

	//判断user表中是否有名字为小１的元组，这里的exist主要是用来指定表的，指定了这里使用的是user表
	has, err = engine.Where("name=?", "小1").Exist(&models.User{})
	fmt.Println(has)

	//判断sql查询有没有对应的结果
	has, err = engine.SQL("select * from user where name = ?", "andrew").Exist()
	fmt.Println(has)

	//判断表存不存在
	has, err = engine.Table("user").Exist()
	fmt.Println(has)

	has, err = engine.Table("user").Where("name=?", "andrew").Exist()
	fmt.Println(has)

	var usr models.User
	ok, err := engine.Table("user").Where("name=?", "andrew").Get(&usr)

	fmt.Println(ok)
	fmt.Println(usr)
}


//从上面的代码中可以知道，Get和Exist非常相似，Get会直接查询出所需要的那一条数据
//而Exist则不会直接查询出数据，而是判断该数据存不存在
//如果你需要判断某一条数据是否存在，如果存在返回该条数据的话，建议直接使用Get
//如果只需要判断一条数据是否存在而不用获取这条数据的话，建议使用Exist
//Exist的效率高于Get
