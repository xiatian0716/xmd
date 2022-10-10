/**
* Copyright (C) 2021-2022
* All rights reserved, Designed By www.github.com/xiatian0716
* 注意：本软件为www.github.com/xiatian0716开发研制
 */
package models

type SysPing struct {
	Name string `json:"name" db:"name"`
}

func (SysPing) TableName() string {
	return "sys_ping"
}

func GetPing() SysPing {
	// db查询
	return SysPing{Name: "sys_ping"}
}
