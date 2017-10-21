package main

import (
	_ "keshihua/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"keshihua/models"
)
func init(){
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("jdbc.username") + ":" + beego.AppConfig.String("jdbc.password") + "@/keshihua?charset=utf8&parseTime=true&charset=utf8&loc=Asia%2FShanghai", 30)
	orm.RegisterModel(new(models.book))
	orm.RunSyncdb("default", false, true)
}
func main() {
	beego.Run()
}

