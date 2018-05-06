package main

import (
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"iris-exercise/orm/xorm/models"
	"fmt"
	"time"
)


func main(){
	engine,err:=xorm.NewEngine("sqlite3","./master.sqlite3")
	if err!=nil{
		fmt.Println("create engine failed!")
		return
	}

	fmt.Println("without transaction:")
	t1_start:=time.Now()
	s1:=engine.NewSession()
	u1:=models.User{Name:"a",Age:10,Passwd:"A"}
	u2:=models.User{Name:"a",Age:10,Passwd:"A"}
	u3:=models.User{Name:"a",Age:10,Passwd:"A"}
	u4:=models.User{Name:"a",Age:10,Passwd:"A"}
	u5:=models.User{Name:"a",Age:10,Passwd:"A"}
	s1.Insert(&u1,&u2,&u3,&u4,&u5)
	count,_:=s1.Count(&models.User{Name:"小1"})
	fmt.Println(count)
	s1.Close()
	delta_t1:=time.Now().Sub(t1_start).Seconds()
	fmt.Println("without transaction spend %f seconds",delta_t1)

	fmt.Println("with transaction:")
	t2_start:=time.Now()
	s2:=engine.NewSession()
	s2.Begin()
	u6:=models.User{Name:"a",Age:10,Passwd:"A"}
	u7:=models.User{Name:"a",Age:10,Passwd:"A"}
	u8:=models.User{Name:"a",Age:10,Passwd:"A"}
	u9:=models.User{Name:"a",Age:10,Passwd:"A"}
	u10:=models.User{Name:"a",Age:10,Passwd:"A"}
	s2.Insert(&u7,&u8,&u9,&u10,&u6)
	count2,_:=s1.Count(&models.User{Name:"小1"})
	fmt.Println(count2)
	s2.Commit()
	s2.Close()
	delta_t2:=time.Now().Sub(t2_start).Seconds()
	fmt.Println("with transaction spend %f seconds",delta_t2)
}

/**

Output of this programe:
without transaction:
7
without transaction spend %f seconds 0.605384967
with transaction:
7
with transaction spend %f seconds 0.123141343

可见事务可以让效率更高
 */