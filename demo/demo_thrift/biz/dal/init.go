package dal

import (
	"github.com/showyquasar88/go-mall/demo/demo_thrift/biz/dal/mysql"
	"github.com/showyquasar88/go-mall/demo/demo_thrift/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
