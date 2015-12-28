package main

import (
	"github.com/astaxie/beego"
	"github.com/beego/samples/todo/controllers"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
orm.RegisterDriver("mysql", orm.DR_MySQL)
orm.RegisterDataBase("default", "mysql", "root:@/test?charset=utf8")
}

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/task/", &controllers.TaskController{}, "get:ListTasks;post:NewTask")
	beego.Router("/task/:id:int", &controllers.TaskController{}, "get:GetTask;put:UpdateTask")

	beego.Router("/post/", &controllers.PostController{}, "get:GetAll;post:Post")
	beego.Router("/post/:id:int", &controllers.PostController{}, "get:GetOne;put:Put;delete:Delete")
	beego.Run()
}
