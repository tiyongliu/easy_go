package controllers

import (
	"easy_go/admin/common"
	"easy_go/admin/servers"
	"easy_go/admin/transform"
	"easy_go/md5"

	"github.com/astaxie/beego/logs"
	// "time"
)

type LoginController struct {
	common.BaseController
}

func (c *LoginController) Get() {
	c.Layout = "layout/login-registerLayout.html"
	c.TplName = "pages/login-img.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["script"] = "script/login_register.html"

	// c.LayoutSections["style"] = "style/login_img.html"

	// c.LayoutSections["script"] = "script/login_img.html"
}

func (c *LoginController) HandleLogin() {
	username := c.GetString("username")
	password := c.GetString("password")
	_ = c.GetStrings("checkbox")
	msg, err := common.Unmarshal(&c.Controller)
	username, _ = transform.InterToString(msg["username"])
	username, _ = transform.InterToString(msg["password"])
	if username == "" || len(username) < 6 || password == "" || len(password) < 6 {
		c.History("账号或密码不合法", "")
		return
	}
	// 加密后的密码
	processPwd, err := servers.SelectUserMd5Pwd(username, md5.Md5(password, common.SECRET_KEY))
	if err != nil {
		logs.Alert("用户:"+username+"加密失败", err.Error())
		c.History("账号或密码不合法", "")
		return
	}
	// 跟数据库的比对
	pwd, err := servers.SelectUserMd5Pwd(username, processPwd)
	if err != nil {
		logs.Alert("用户:"+username+"比对密码出多", err.Error())
		c.History("账号或密码不合法", "")
		return
	}
	if pwd != "" {
		c.Redirect("/workplace", 302)
		return
	} else {
		c.Redirect("/login", 302)
	}
}
