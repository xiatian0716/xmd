/**
* Copyright (C) 2021-2022
* All rights reserved, Designed By www.github.com/xiatian0716
* 注意：本软件为www.github.com/xiatian0716开发研制
 */
package vo

type ResultList struct {
	Content       interface{} `json:"content"`
	TotalElements int64       `json:"totalElements"`
	ExtendData    interface{} `json:"extendData"`
}
