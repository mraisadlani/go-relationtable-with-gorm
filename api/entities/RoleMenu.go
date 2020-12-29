package entities

import "time"

type RoleMenu struct {
	IdRoleMenu int64     `json:"id_role_menu" gorm:"column:id_role_menu;PRIMARY_KEY;NOT NULL;AUTO_INCREMENT"`
	IdRole     int64     `json:"id_role" gorm:"column:id_role"`
	IdType     int64     `json:"id_type" gorm:"column:id_type"`
	IdMenu     int64     `json:"id_menu" gorm:"column:id_menu"`
	CreateBy   string    `json:"createby"`
	CreateDate time.Time `json:"createdate" default:CURRENT_TIMESTAMP`
	UpdateBy   string    `json:"updateby"`
	UpdateDate time.Time `json:"updatedate" default:CURRENT_TIMESTAMP`
	Role       Role      `gorm:"foreignKey:IdRole;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Type       Tipe      `gorm:"foreignKey:IdType;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Menu       Menu      `gorm:"foreignKey:IdMenu;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (r *RoleMenu) TableName() string {
	return "tb_role_menu"
}
