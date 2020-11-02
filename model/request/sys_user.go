package request

// user register
type RegisterStruct struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Nickname    string `json:"nickname" gorm:"default:'QMPlusUser'"`
	HeaderImg   string `json:"headerImg" gorm:"default:'http://www.henrongyi.top/avatar/lufu.jpg'"`
	AuthorityId string `json:"authorityId" gorm:"default:888"`
}
