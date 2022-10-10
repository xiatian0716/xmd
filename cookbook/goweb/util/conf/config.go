/**
* Copyright (C) 2021-2022
* All rights reserved, Designed By www.github.com/xiatian0716
* 注意：本软件为www.github.com/xiatian0716开发研制
 */
package conf

import (
	"github.com/xiatian0716/xmd/goweb/util/mysqlx"
	"github.com/xiatian0716/xmd/goweb/util/redisxx"
	"github.com/xiatian0716/xmd/goweb/util/snowflakex"
	"github.com/xiatian0716/xmd/goweb/util/zapx"
)

// Conf 全局变量，用来保存程序的所有配置信息
var Conf = new(Config)

type Config struct {
	App       App                  `mapstructure:"app"`
	Snowflake snowflakex.Snowflake `mapstructure:"snowflake"`
	MySQL     mysqlx.MySQL         `mapstructure:"mysql"`
	Redis     redisxx.Redis        `mapstructure:"redis"`
	Log       zapx.Log             `mapstructure:"log"`
}

type App struct {
	Name     string `mapstructure:"name"`
	Mode     string `mapstructure:"mode"`
	Version  string `mapstructure:"version"`
	Port     int    `mapstructure:"port"`
	PageSize int    `mapstructure:"page-size"`
}

//type Snowflakex struct {
//	StartTime string `mapstructure:"start_time"`
//	MachineID int64  `mapstructure:"machine_id"`
//}

//type MySQL struct {
//	Host         string `mapstructure:"host"`
//	User         string `mapstructure:"user"`
//	Password     string `mapstructure:"password"`
//	DB           string `mapstructure:"dbname"`
//	Port         int    `mapstructure:"port"`
//	MaxOpenConns int    `mapstructure:"max_open_conns"`
//	MaxIdleConns int    `mapstructure:"max_idle_conns"`
//}

//type Redis struct {
//	Host        string        `mapstructure:"host"`
//	Password    string        `mapstructure:"password"`
//	MaxIdle     int           `mapstructure:"max-idle"`
//	MaxActive   int           `mapstructure:"max-active"`
//	IdleTimeout time.Duration `mapstructure:"idle-timeout"`
//}

//type Log struct {
//	Level      string `mapstructure:"level"`
//	Filename   string `mapstructure:"filename"`
//	MaxSize    int    `mapstructure:"max_size"`
//	MaxAge     int    `mapstructure:"max_age"`
//	MaxBackups int    `mapstructure:"max_backups"`
//}
