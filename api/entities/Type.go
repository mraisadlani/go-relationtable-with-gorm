package entities

import (
	"time"
)

type Tipe struct {
	Id         int64     `json:"id_type" gorm:"column:id_type;PRIMARY_KEY;NOT NULL;AUTO_INCREMENT"`
	Type       string    `json:"type"`
	Keterangan string    `json:"keterangan"`
	Status     string    `json:"status"`
	CreateBy   string    `json:"createby"`
	CreateDate time.Time `json:"createdate" default:CURRENT_TIMESTAMP`
	UpdateBy   string    `json:"updateby"`
	UpdateDate time.Time `json:"updatedate" default:CURRENT_TIMESTAMP`
}

func (t *Tipe) TableName() string {
	return "tb_type"
}
