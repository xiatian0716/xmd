/**
* Copyright (C) 2021-2022
* All rights reserved, Designed By www.github.com/xiatian0716
* 注意：本软件为www.github.com/xiatian0716开发研制
 */
package app

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	err := c.Bind(form)
	if err != nil {
		zap.L().Error(err.Error())
		return http.StatusBadRequest, INVALID_PARAMS
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		zap.L().Error(err.Error())
		return http.StatusInternalServerError, ERROR
	}
	if !check {
		MarkErrors(valid.Errors)
		return http.StatusBadRequest, INVALID_PARAMS
	}

	return http.StatusOK, SUCCESS
}

func BindAndValidate(c *gin.Context, form interface{}) error {
	err := c.Bind(form)
	if err != nil {
		zap.L().Error(err.Error())
		return err
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		zap.L().Error(err.Error())
		return err
	}
	if !check {
		MarkErrors(valid.Errors)
		return buildFormErr(valid.Errors)
		//return http.StatusBadRequest, constant.INVALID_PARAMS
	}

	return nil
}

func buildFormErr(errs []*validation.Error) error {
	var msg strings.Builder
	for _, v := range errs {
		if v.Field != "" {
			msg.WriteString(v.Field)
		} else if v.Key != "" {
			msg.WriteString(v.Key)
		} else {
			msg.WriteString(v.Name)
		}

		msg.WriteString(" : ")
		if v.Value != nil {
			b, _ := json.Marshal(v.Value)
			msg.WriteString(string(b))
		}

		msg.WriteString(" => ")
		msg.WriteString(v.Error())
		//msg.WriteString(" should=> ")
		//bb,_ := json.Marshal(v.LimitValue)
		//msg.WriteString(string(bb))
	}
	return errors.New(msg.String())
}
