package controllers

import (
	"easy_go/blog/logger"
	"easy_go/common"
	"easy_go/lib"
	"encoding/json"
	"gitee.com/zchunshan/d3auth"
	"github.com/astaxie/beego"
)

type OAuthControllers struct {
	common.BaseController
}

type Auth_conf struct {
	Appid  string
	Appkey string
	Rurl   string
}

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"` // 这个字段没用到
	Scope       string `json:"scope"`      // 这个字段也没用到
}

func (c *OAuthControllers) Get() {
	codeStr := c.GetString("code")
	user, err := github(codeStr)

	logger.Info(user, err)
	if err != nil {
		c.Error("第三方登录失败")
		return
	}

	beego.Info(c.Ctx.Request.RequestURI)
	c.Success(user)
}

func github(codeStr string) (map[string]interface{}, error) {

	github_conf := &d3auth.Auth_conf{Appid: lib.Conf.Read("github", "ClientId"), Appkey: lib.Conf.Read("github", "ClientSecret"), Rurl: lib.Conf.Read("github", "RedirectUrl")}
	githubAuth := d3auth.NewAuth_github(github_conf)
	token, err := githubAuth.Get_Token(codeStr)

	if err != nil {
		logger.Info("获取github登录token失败", err.Error())
		return nil, err
	}

	msg, err := githubAuth.Get_User_Info(token)
	if err != nil {
		logger.Info("获取github用户信息失败", err.Error())
		return nil, err
	}

	var userInfo map[string]interface{}

	if err := json.Unmarshal([]byte(msg), &userInfo); err != nil {
		logger.Info("github登录用户信息转json失败", err.Error())
		return nil, err
	}

	userInfo["access_token"] = token

	return userInfo, nil
}
