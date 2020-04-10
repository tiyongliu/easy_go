package routers

import (
	"easy_go/admin/common"
	"easy_go/admin/controllers"
	"easy_go/admin/controllers/article"
	"github.com/astaxie/beego/context"

	"github.com/astaxie/beego"
)

const Api = "/api"

func init() {
	beego.InsertFilter("/menuSetting/add", beego.BeforeRouter, FilterUser)
	// beego.Router("/", &controllers.IndexControllers{}) // 废弃
	beego.Router("/login", &controllers.LoginController{})

	beego.Router("/register", &controllers.RegisterController{})

	beego.Router("/", &controllers.DashBoardControllers{})

	beego.Router("/analysis", &controllers.DashBoardControllers{}, "get:HandleAnalysis")

	// 工作台
	beego.Router("/workplace", &controllers.DashBoardControllers{}, "get:HandleWorkplace")

	// 路由权限设置
	beego.Router("/menuSetting", &controllers.MenuController{})

	// 导航菜单-add
	beego.Router("/menuSetting/add", &controllers.MenuController{}, "get:Add;post:HandleMenuAdd")

	// 导航菜单-info
	beego.Router("/menuSetting/info", &controllers.MenuController{}, "get:Info")

	// 文章类型
	beego.Router("/article/type", &article.ArticleControllerType{})

	// 文章列表
	beego.Router("/article/list", &article.ArticleList{})

	// 文章新增+编辑
	beego.Router("/article/details", &article.ArticleDetails{}, "get:AddOfUpdate")
	register()

}

func register() {
	beego.Router(Api+"/login", &controllers.LoginController{}, "post:HandleLogin")
	beego.Router(Api+"/register", &controllers.RegisterController{}, "post:AddRegister")
}

// 全局过滤方法。
// https://www.kancloud.cn/hello123/beego/126127
var FilterUser = func(ctx *context.Context) {

	_, ok := ctx.Input.Session("userName").(string)
	if !ok && ctx.Request.RequestURI != "/login" || ctx.Request.RequestURI != "register" {
		// 获取cookies
		cook := ctx.GetCookie("auth")
		if cook == "" {
			ctx.Redirect(302, "/login")
		} else {
			common.ParseTokenUser(cook)
		}
	}
}
