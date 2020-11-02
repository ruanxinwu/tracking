package initialize

import (
	"tracking/global"
)
import "tracking/model"

// 注册数据库表专用
func DBTables() {
	global.GVA_DB.AutoMigrate(model.JwtBlackList{})

	global.GVA_LOG.Debug("register table success")
}
