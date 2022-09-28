package redisxx

import "time"

type Redis struct {
	Host        string        `mapstructure:"host"`
	Password    string        `mapstructure:"password"`
	MaxIdle     int           `mapstructure:"max-idle"`
	MaxActive   int           `mapstructure:"max-active"`
	IdleTimeout time.Duration `mapstructure:"idle-timeout"`
}
