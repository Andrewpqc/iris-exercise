/*
名称映射规则主要负责结构体名称到表名和结构体field到表字段的名称映射.
由core.IMapper接口的实现者来管理，xorm内置了三种IMapper实现：core.SnakeMapper ，
core.SameMapper和core.GonicMapper。
* SnakeMapper 支持struct为驼峰式命名，表结构为下划线命名之间的转换，这个是默认的Maper；
* SameMapper 支持结构体名称和对应的表名称以及结构体field名称与对应的表字段名称相同的命名；
* GonicMapper 和SnakeMapper很类似，但是对于特定词支持更好，比如ID会翻译成id而不是i_d。

当前SnakeMapper为默认值，如果需要改变时，在engine创建完成后使用
engine.SetMapper(core.SameMapper{})
表名称和字段名称的映射规则默认是相同的，当然也可以设置为不同，如：
engine.SetTableMapper(core.SameMapper{})
engine.SetColumnMapper(core.SnakeMapper{})

前缀映射，后缀映射和缓存映射
tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "prefix_")
engine.SetTableMapper(tbMapper)

core.NewSufffixMapper(core.SnakeMapper{}, "suffix")

core.NewCacheMapper(core.SnakeMapper{})



使用Table和Tag改变名称映射
*/

package main

import (
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
	_ "github.com/mattn/go-sqlite3"
	//"iris-exercise/orm/xorm/models"
	"fmt"
)


func main(){
	engine,err:=xorm.NewEngine("sqlite3","./master.sqlite3")
	if err!=nil{
		fmt.Println("create engine failed!")
		return
	}
	//驼峰－>下划线命名规则
	engine.SetMapper(core.SnakeMapper{})
	//和上面的规则相似，只是对特定的词支持更好
	engine.SetMapper(core.GonicMapper{})
	//结构名与数据库各项的名字相同
	engine.SetMapper(core.SameMapper{})

	//也可单独设置表名的映射形式和字段名的映射形式
	engine.SetTableMapper(core.SameMapper{})
	engine.SetColumnMapper(core.GonicMapper{})
}