package main

import (
	"tracking/core"
	"tracking/global"
	"tracking/initialize"
)

func main() {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		initialize.Mysql()
	default:
		initialize.Mysql()
	}

	initialize.DBTables()

	// 结束关闭数据库链接
	defer global.GVA_DB.Close()
	core.RunWindowsServer()
}
