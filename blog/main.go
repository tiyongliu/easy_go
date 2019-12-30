package main

import (
	"easy_go/blog/controllers"
	_ "easy_go/blog/routers"
	"github.com/astaxie/beego"
)


func main() {
	// beego.BConfig.RunMode = beego.AppConfig.String("runmode")          // 环境
	beego.ErrorController(&controllers.ErrorController{})
    beego.BConfig.RunMode = "dev"
	beego.BConfig.WebConfig.ViewsPath = "template" // 静态目录
	beego.Run(":" + beego.AppConfig.String("httpport"))
}
