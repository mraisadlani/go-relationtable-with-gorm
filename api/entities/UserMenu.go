package entities

import "time"

type UserMenu struct {
	IdUserMenu int64     `json:"id_user_menu" gorm:"column:id_user_menu;PRIMARY_KEY;NOT NULL;AUTO_INCREMENT"`
	IdUser     int64     `json:"id_user" gorm:"column:id_user"`
	IdMenu     int64     `json:"id_menu" gorm:"column:id_menu"`
	CreateBy   string    `json:"createby"`
	CreateDate time.Time `json:"createdate" default:CURRENT_TIMESTAMP`
	UpdateBy   string    `json:"updateby"`
	UpdateDate time.Time `json:"updatedate" default:CURRENT_TIMESTAMP`
	User       User      `gorm:"foreignKey:IdUser;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Menu       Menu      `gorm:"foreignKey:IdMenu;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (u *UserMenu) TableName() string {
	return "tb_user_menu"
}
