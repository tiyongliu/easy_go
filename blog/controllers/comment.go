package controllers

import (
	"easy_go/blog/logger"
	"easy_go/blog/servers"
	"easy_go/common"
	myjwt "easy_go/middleware"
	"easy_go/transform"
	"strconv"
)

type CommentControllers struct {
	common.BaseController
}

// 查询评论信息
func (c *CommentControllers) GetCommentList() {
	param := c.Ctx.Input.Param(":id")
	if param == "" {
		c.Error("获取评论参数不合法")
		return
	}

	_id, err := strconv.Atoi(param)
	if err != nil || _id <= 0 {
		c.Error("获取评论参数不合法")
		return
	}

	pageStr := c.GetString("page")
	if pageStr == "" {
		c.Error("页码不能为空")
		return
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		c.Error("页码不合法")
		return
	}

	// 获取到文章id去查询评论+回复
	cl, err := servers.SelectCommentList(_id, common.PAGE_SIZE, page)
	if err != nil {
		c.Error("查询数据失败")
		return
	}

	c.Success(cl)
}

// 新增评论
func (c *CommentControllers) InsertComment() {
	auth := c.Ctx.Request.Header.Get("r")
	j := myjwt.NewJWT()
	claims, err := j.ParseToken(auth)
	if err != nil {
		logger.Info("评论解析token失败", err.Error())
		c.Error("评论失败，参数不合法！1")
		return
	}

	role, err := servers.Select_github(claims.ID, claims.Username, claims.LoginIp, auth)
	if err != nil {
		c.Error("评论失败，参数不合法！2")
		return
	}

	msg, err := common.Unmarshal(&c.Controller)
	if err != nil {
		c.Error("评论失败，参数不合法！3")
		return
	}

	message, err := transform.InterToString(msg["message"])
	if err != nil && message == "" {
		c.Error("评论失败！4")
		return
	}

	articleId, err := transform.InterToInt(msg["article_id"])

	// 新增评论信息
	err = servers.AddComment(role, articleId, message)
	if err != nil {
		c.Error("评论失败！")
		return
	}

	c.Success(articleId)
}
