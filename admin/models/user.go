package models

import (
	"database/sql"
	"time"
)

type User struct {
	Id               int          `json:"id"`
	UserName         string       `json:"user_name"`
	PassWord         string       `json:"pass_word"`
	Role             int          `json:"role"`               // 用户角色，暂时这个字段没有用。
	LoginIp          string       `json:"login_ip"`           // 用户登陆的ip
	CurrentLoginTime sql.NullTime `json:"current_login_time"` // 登陆时间
	CreatedTime      time.Time    `json:"created_time"`       // 创建时间
	UpdateTime       sql.NullTime `json:"update_time"`        // 更新时间
	AuthToken        string       `json:"auth_token"`         // token
}

type RoleTypeEle int8

const (
	GITHUB_ ReplyTypeEle = 1
	QQ_     ReplyTypeEle = 2
)

// token 第三方登陆的65b5871477065bb229fda33bc22f2f6befad5be0
type Role struct {
	Id               int          `json:"id"`
	Name             string       `json:"name"`               // 用户名
	UserType         RoleTypeEle  `json:"user_type"`          // 用户类型 1：github 2：qq
	UserId           int          `json:"user_id"`            // 用户id
	UserLogin        string       `json:"user_login"`         // 用户登陆
	Location         string       `json:"location"`           // 位置
	AvatarUrl        string       `json:"avatar_url"`         // 头像url
	LoginIp          string       `json:"login_ip"`           // 用户登陆的ip
	CurrentLoginTime sql.NullTime `json:"current_login_time"` // 登陆时间
	CreatedTime      time.Time    `json:"created_time"`       // 创建时间
	UpdateTime       sql.NullTime `json:"update_time"`        // 更新时间
	AuthToken        string       `json:"auth_token"`         // token
}
