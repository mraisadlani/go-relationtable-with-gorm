package entities

import "time"

type Role struct {
	Id         int64      `json:"id_role" gorm:"column:id_role;PRIMARY_KEY;NOT NULL;AUTO_INCREMENT"`
	Role       string     `json:"role"`
	Status     int64      `json:"status"`
	CreateBy   string     `json:"createby"`
	CreateDate time.Time  `json:"createdate" default:CURRENT_TIMESTAMP`
	UpdateBy   string     `json:"updateby"`
	UpdateDate time.Time  `json:"updatedate" default:CURRENT_TIMESTAMP`
	RoleUsers  []RoleUser `gorm:"ForeignKey:IdRole"`
	RoleMenus  []RoleMenu `gorm:"ForeignKey:IdRole"`
}

func (r *Role) TableName() string {
	return "tb_role"
}
