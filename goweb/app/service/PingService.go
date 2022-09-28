/**
* Copyright (C) 2021-2022
* All rights reserved, Designed By www.github.com/xiatian0716
* 注意：本软件为www.github.com/xiatian0716开发研制
 */
package service

import (
	"github.com/xiatian0716/xmd/goweb/app/models"
	"github.com/xiatian0716/xmd/goweb/util/app/vo"
)

type Ping struct {
	Id        int64
	Name      string
	Something string
}

func (p *Ping) GetPings() vo.ResultList {
	pings := models.GetPing()
	return vo.ResultList{Content: pings, TotalElements: 0}
}
