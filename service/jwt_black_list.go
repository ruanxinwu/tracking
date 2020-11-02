package service

import (
	"tracking/global"
	"tracking/model"
)

func IsBlackList(jwt string, JwtList model.JwtBlackList) bool {
	return !global.GVA_DB.Where("jwt = ?", jwt).First(&JwtList).RecordNotFound()
}
