/**
* Copyright (C) 2021-2022
* All rights reserved, Designed By www.github.com/xiatian0716
* 注意：本软件为www.github.com/xiatian0716开发研制
 */
package dto

type AuthUser struct {
	Code     string `json:"code"`
	Password string `json:"password"`
	Username string `json:"username"`
	Id       string `json:"uuid"`
}
