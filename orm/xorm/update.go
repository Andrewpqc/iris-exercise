//Update 更新数据，除非使用Cols,AllCols函数指明，默认只更新非空和非0的字段

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

	var user = models.User{Name: "updateName"}

	//将ID为1的元组的Name字段改为updateName
	//affected, err := engine.ID(1).Update(&user)
	//fmt.Println(affected)
	// UPDATE user SET ... Where id = ?

	//将名字为andrew的元组的名字改为updateName
	//affected, err := engine.Update(&user, &models.User{Name:"andrew"})
	//fmt.Println(affected)
	// UPDATE user SET ... Where name = ?

	//将Id为1或2或3的用户用户的Name字段改为updateName
	var ids = []int64{1, 2, 3}
	affected, err := engine.In("id", ids).Update(&user)
	fmt.Println(affected)
	// UPDATE user SET ... Where id IN (?, ?, ?)

	//只更新了age列的值,Name没有被更新
	// force update indicated columns by Cols
	affected, err = engine.ID(1).Cols("age").Update(&models.User{Name: "newName", Age: 2222})
	fmt.Println(affected)
	// UPDATE user SET age = ?, updated=? Where id = ?

	//指定不更新某一列
	// force NOT update indicated columns by Omit
	affected, err = engine.ID(1).Omit("name").Update(&models.User{Name: "name", Age: 12})
	fmt.Println(affected)
	// UPDATE user SET age = ?, updated=? Where id = ?

	//所有的类均可以更新
	affected, err = engine.ID(1).AllCols().Update(&models.User{Name: "hhhh", Age: 1555})
	fmt.Println(affected)
	// UPDATE user SET name=?,age=?,salt=?,passwd=?,updated=? Where id = ?

}
