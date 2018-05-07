package main

import (
	"github.com/jinzhu/gorm"
	"iris-exercise/orm/gorm/models"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"fmt"
)

func main(){
	db,err:=gorm.Open("sqlite3","test1.sqlite3")
	if err!=nil{
		panic("failed to connect database")
	}
	defer db.Close()

	//提供自动迁移支持
	db.AutoMigrate(&models.Product{})

	//创建了一条数据
	//db.Create(&models.Product{Code:"L1212",Price:1000})
	//db.Create(&models.Product{Code:"L1213",Price:1200})
	var p1,p2 models.Product
	db.First(&p1,2)//find product with id 1
	fmt.Println(p1)
	db.First(&p2,"code=?","L1212")//find product with code L1212
	fmt.Println(p2)

	//更新数据
	//db.Model(&p1).Update("Price",2000)
	db.Delete(&p1)



}