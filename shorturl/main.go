package main

import (
	"github.com/astaxie/beego"
	"github.com/beego/samples/shorturl/controllers"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/v1/shorten", &controllers.ShortController{})
	beego.Router("/v1/expand", &controllers.ExpandController{})
	beego.Run()
}
