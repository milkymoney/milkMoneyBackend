package main

import (
	"fmt"
	_ "apiproject/routers"
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
)

type Test struct {
    Id   int
    Name string `orm:"size(100)"`
}

func init() {
    orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "wty:97112500@tcp(127.0.0.1:3306)/test?charset=utf8")
	orm.RegisterModel(new(Test))
	orm.RunSyncdb("default", false, true)
}



func main() {
	o := orm.NewOrm()
	test := Test{Name:"test"}
	
	//insert
	id, err := o.Insert(&test)
	fmt.Printf("ID:%d, ERR:%v\n",id,err)

	//update
	test.Name = "noTest"
	num,err := o.Update(&test)
	fmt.Printf("NUM:%d, ERR=%v\n",num,err)

	//read one
	u := Test{Id: test.Id}
	err = o.Read(&u)

	//delete
	num, err = o.Delete(&u)
	fmt.Printf("Num:%d, ERR:%v\n",num,err)

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
