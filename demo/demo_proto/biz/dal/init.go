package dal

import (
	"github.com/showyquasar88/go-mall/demo/demo_proto/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
