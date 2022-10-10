/**
* Copyright (C) 2021-2022
* All rights reserved, Designed By www.github.com/xiatian0716
* 注意：本软件为www.github.com/xiatian0716开发研制
 */
package vo

type JwtUser struct {
	Id       int64
	Avatar   string
	Email    string
	Username string
	Phone    string
	NickName string
	Sex      string
	Dept     string
	Job      string
	Roles    []string
}
