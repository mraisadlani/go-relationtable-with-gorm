package entities

import "time"

type Menu struct {
	Id         int64      `json:"id_menu" gorm:"column:id_menu;PRIMARY_KEY;NOT NULL;AUTO_INCREMENT"`
	Menu       string     `json:"menu"`
	Icon       string     `json:"icon"`
	Path       string     `json:"path"`
	SubMenu    []SubMenu  `json:"child_menu" gorm:"foreignKey:IdMenu;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Status     int64      `json:"status"`
	CreateBy   string     `json:"createby"`
	CreateDate time.Time  `json:"createdate" default:CURRENT_TIMESTAMP`
	UpdateBy   string     `json:"updateby"`
	UpdateDate time.Time  `json:"updatedate" default:CURRENT_TIMESTAMP`
	UserMenus  []UserMenu `gorm:"ForeignKey:IdMenu"`
}

func (m *Menu) TableName() string {
	return "tb_menu"
}
