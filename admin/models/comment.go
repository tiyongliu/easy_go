package models

import (
	"database/sql"
	"time"
)

// 网站点赞 评论 回复 数据库设计
// https://blog.csdn.net/lm1622/article/details/77950133?utm_medium=distribute.pc_relevant_t0.none-task-blog-BlogCommendFromMachineLearnPai2-1.edu_weight&depth_1-utm_source=distribute.pc_relevant_t0.none-task-blog-BlogCommendFromMachineLearnPai2-1.edu_weight
// https://blog.csdn.net/ztchun/article/details/71106117
// https://www.jianshu.com/p/f9e27a96da89

/*评论*/
type Comment struct {
	Id           int          `json:"id"`
	ArticleId    int          `json:"article_id"`    // 文章id
	Content      string       `json:"content"`       // 评论内容
	UserId       int          `json:"user_id"`       // 评论用户id
	CommentState bool         `json:"comment_state"` // 评论状态：默认显示全部，超级管理员可以删除评论
	CreatedTime  time.Time    `json:"created_time"`  // 创建时间
	UpdateTime   sql.NullTime `json:"update_time"`   // 更新时间
}

type ReplyTypeEle int8

const (
	Comment_ ReplyTypeEle = 1
	Reply_   ReplyTypeEle = 2
)

// 回复表
type Reply struct {
	Id        int          `json:"id"`
	CommentId int          `json:"comment_id"`               // 评论id
	Content   string       `json:"content"`                  // 回复内容
	UserId    int          `json:"user_id"`                  // 回复用户id
	ReplyType ReplyTypeEle `json:"reply_type" gorm:"size:8"` // 表示回复的类型，因为回复可以是针对评论的回复(comment表)，也可以是针对回复的回复(reply表)， 通过这个字段来区分两种情景。
	ReplyId   int          `json:"reply_id"`                 // 表示回复目标的id，如果reply_type是comment的话，那么reply_id＝commit_id，如果reply_type是reply的话，这表示这条回复的父回复
}

type ZanTypeEle int8

const (
	ArticleZan_ ZanTypeEle = 1
	CommentZan_ ZanTypeEle = 2
	ReplyZan_   ZanTypeEle = 3
)

// 点赞表
type Zan struct {
	Id     int        `json:"id"`
	TypeId int        `json:"type_id"` // 对应的作品或评论的id
	Type   ZanTypeEle `json:"type"`    // 点赞类型  1作品点赞  2 评论点赞 3  回复点赞
	UserId int        `json:"user_id"` // 用户id
	State  bool       `json:"state"`   // 点赞状态  0--取消赞   1--有效赞
}
