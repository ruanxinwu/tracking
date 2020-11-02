package core

import (
	"tracking/global"
	"tracking/initialize"
)

func RunWindowsServer() {
	if global.GVA_CONFIG.System.UseMultipoint {
		initialize.Redis()
	}

	_ = initialize.Routers()
}
