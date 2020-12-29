package entities

import "time"

type SubMenu struct {
	Id         int64     `json:"id_sub_menu" gorm:"column:id_sub_menu;PRIMARY_KEY;NOT NULL;AUTO_INCREMENT"`
	SubMenu    string    `json:"sub_menu"`
	IdMenu     int64     `json:"id_menu" gorm:"column:id_menu"`
	Icon       string    `json:"icon"`
	Path       string    `json:"path"`
	Status     int64     `json:"status"`
	CreateBy   string    `json:"createby"`
	CreateDate time.Time `json:"createdate" default:CURRENT_TIMESTAMP`
	UpdateBy   string    `json:"updateby"`
	UpdateDate time.Time `json:"updatedate" default:CURRENT_TIMESTAMP`
}

func (s *SubMenu) TableName() string {
	return "tb_sub_menu"
}
