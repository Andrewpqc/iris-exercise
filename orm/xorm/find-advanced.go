package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/go-xorm/xorm"
	"iris-exercise/orm/xorm/models"
	"fmt"
	"time"
)

func main(){
	engine,err:=xorm.NewEngine("sqlite3","./find-advanced.sqlite3")
	if err!=nil{
		fmt.Println("create engine failed!")
		return
	}


	//注意，这里的UserDetail是一个虚表，数据库中不会存在对应的表，这里知识作为获取数据的结构
	engine.Sync2(new(models.User),new(models.Item),new(models.Detail),new(models.UserDetail))
	u1:=models.User{Name:"1",Age:1,Passwd:"1",Created:time.Now(),Updated:time.Now()}
	u2:=models.User{Name:"2",Age:2,Passwd:"2",Created:time.Now(),Updated:time.Now()}
	u3:=models.User{Name:"3",Age:3,Passwd:"3",Created:time.Now(),Updated:time.Now()}
	engine.Insert(&u1,&u2,&u3)

	d1:=models.Detail{Job:"IT",UserId:1}
	d2:=models.Detail{Job:"Edu",UserId:1}
	d3:=models.Detail{Job:"IT",UserId:3}
	d4:=models.Detail{Job:"IT",UserId:2}
	engine.Insert(&d1,&d2,&d3,&d4)

	//联结操作
	var users []models.UserDetail//跨表获取信息
	engine.Table("user").Select("user.*, detail.*").
	 	Join("INNER", "detail", "detail.user_id = user.id").
		Where("user.name = ?", "1l").Limit(10, 0).
		Find(&users)
	for _,user:=range users{
		fmt.Println(user)
	}

}