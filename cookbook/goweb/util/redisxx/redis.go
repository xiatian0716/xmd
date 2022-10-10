/**
* Copyright (C) 2021-2022
* All rights reserved, Designed By www.github.com/xiatian0716
* 注意：本软件为www.github.com/xiatian0716开发研制
 */
package redisxx

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"time"
)

var RedisConn *redis.Pool

// Setup Initialize the Redis instance
func Setup(r Redis) error {
	RedisConn = &redis.Pool{
		MaxIdle:     r.MaxIdle,
		MaxActive:   r.MaxActive,
		IdleTimeout: r.IdleTimeout * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", r.Host)
			if err != nil {
				return nil, err
			}
			if r.Password != "" {
				if _, err := c.Do("AUTH", r.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return nil
}

func RedisClose() {
	_ = RedisConn.Close()
}

func SetEx(key, val string, exp int64) error {
	conn := RedisConn.Get()
	defer conn.Close()
	_, err := conn.Do("SETEX", key, exp, val)
	return err
}

func GetMap(key string) (map[string]string, error) {
	conn := RedisConn.Get()
	defer conn.Close()
	rsp := make(map[string]string)
	res, err := conn.Do("GET", key)
	rsp[key] = string(res.([]byte))
	return rsp, err
}

// Set a key/value
func Set(key string, data interface{}, time int) error {
	conn := RedisConn.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}
	if time > 0 {
		_, err = conn.Do("EXPIRE", key, time)
		if err != nil {
			return err
		}
	}

	return nil
}

// Exists check a key
func Exists(key string) bool {
	conn := RedisConn.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

// Get get a key
func GetString(key string) string {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return ""
	}

	return reply
}

// Get a key
func Get(key string) ([]byte, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

// Delete delete a kye
func Delete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

// LikeDeletes batch delete
func LikeDeletes(key string) error {
	conn := RedisConn.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}

	for _, key := range keys {
		_, err = Delete(key)
		if err != nil {
			return err
		}
	}

	return nil
}
