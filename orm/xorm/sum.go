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

	agesFloat64, err := engine.Sum(&user, "age")
	// SELECT sum(age) AS total FROM user

	agesInt64, err := engine.SumInt(&user, "age")
	// SELECT sum(age) AS total FROM user

	sumFloat64Slice, err := engine.Sums(&user, "age", "score")
	// SELECT sum(age), sum(score) FROM user

	sumInt64Slice, err := engine.SumsInt(&user, "age", "score")
	// SELECT sum(age), sum(score) FROM user

}
