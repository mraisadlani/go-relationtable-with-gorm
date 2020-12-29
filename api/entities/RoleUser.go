package entities

import "time"

type RoleUser struct {
	IdRoleUser int64     `json:"id_role_user" gorm:"column:id_role_user;PRIMARY_KEY;AUTO_INCREMENT;NOT NULL;"`
	IdUser     int64     `json:"id_user" gorm:"column:id_user"`
	IdRole     int64     `json:"id_role" gorm:"column:id_role"`
	CreateBy   string    `json:"createby"`
	CreateDate time.Time `json:"createdate" default:CURRENT_TIMESTAMP`
	UpdateBy   string    `json:"updateby"`
	UpdateDate time.Time `json:"updatedate" default:CURRENT_TIMESTAMP`
	User       User      `gorm:"foreignKey:IdUser;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Role       Role      `gorm:"foreignKey:IdRole;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (r *RoleUser) TableName() string {
	return "tb_role_user"
}
