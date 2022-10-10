/**
* Copyright (C) 2021-2022
* All rights reserved, Designed By www.github.com/xiatian0716
* 注意：本软件为www.github.com/xiatian0716开发研制
 */
package app

import (
	"github.com/astaxie/beego/validation"
	"go.uber.org/zap"
)

// MarkErrors logs error logs
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		zap.L().Info(err.Key, zap.String("MarkErrors", err.Message))
	}

	return
}
