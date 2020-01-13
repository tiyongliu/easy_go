package routers

import (
	"easy_go/admin/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// beego.Router("/", &controllers.IndexControllers{}) // 废弃

	beego.Router("/login", &controllers.LoginController{})

	beego.Router("/register", &controllers.RegisterController{})

	beego.Router("/welcome", &controllers.WelcomeController{})

	beego.Router("/", &controllers.MainController{})
}
