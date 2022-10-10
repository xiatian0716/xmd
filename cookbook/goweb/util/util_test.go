package util

import (
	"database/sql"
	"fmt"
	"github.com/xiatian0716/xmd/goweb/util/conf"
	"github.com/xiatian0716/xmd/goweb/util/mysqlx"
	"github.com/xiatian0716/xmd/goweb/util/redisxx"
	"testing"
	"time"
)

type SysJob struct {
	Id         int32        `db:"id"`
	Name       string       `db:"name"`
	Enabled    bool         `db:"enabled"`
	Sort       int16        `db:"sort"`
	DeptId     int8         `db:"dept_id"`
	CreateTime time.Time    `db:"create_time"`
	UpdateTime sql.NullTime `db:"update_time"`
	IsDel      bool         `db:"is_del"`
}

func TestMysqlConn(t *testing.T) {
	_ = conf.Setup(".")
	_ = mysqlx.Setup(conf.Conf.MySQL)
	defer mysqlx.MysqlClose()

	sqlStr := `select * from sys_job`
	var job []SysJob
	err := mysqlx.Db.Select(&job, sqlStr)
	fmt.Println(err)
	fmt.Println(job)
}

func TestRedisConn(t *testing.T) {
	_ = conf.Setup(".")
	_ = redisxx.Setup(conf.Conf.Redis)
	defer redisxx.RedisClose()
}
