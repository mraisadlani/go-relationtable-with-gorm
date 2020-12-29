package entities

import "time"

type User struct {
	Id          int64      `json:"id_user" gorm:"PRIMARY_KEY;NOT NULL;AUTO_INCREMENT"`
	Username    string     `json:"username"`
	Password    string     `json:"password"`
	NamaLengkap string     `json:"nama_lengkap"`
	Status      int64      `json:"status"`
	Email       string     `json:"email"`
	Photo       string     `json:"photo"`
	TypeId      int64      `json:"id_type" gorm:"column:id_type"`
	Type        Tipe       `gorm:"foreignkey:TypeId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreateBy    string     `json:"createby"`
	CreateDate  time.Time  `json:"createdate" default:CURRENT_TIMESTAMP`
	UpdateBy    string     `json:"updateby"`
	UpdateDate  time.Time  `json:"updatedate" default:CURRENT_TIMESTAMP`
	RoleUsers   []RoleUser `gorm:"ForeignKey:IdUser"`
	UserMenus   []UserMenu `gorm:"ForeignKey:IdUser"`
}

func (u *User) TableName() string {
	return "tb_user"
}
