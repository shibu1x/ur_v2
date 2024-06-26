// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameRoom = "rooms"

// Room mapped from table <rooms>
type Room struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	HouseCode string    `gorm:"column:house_code;not null" json:"house_code"`
	RoomCode  string    `gorm:"column:room_code;not null" json:"room_code"`
	Status    string    `gorm:"column:status;not null;default:ready" json:"status"`
	GotAt     time.Time `gorm:"column:got_at;not null;default:1000-01-01 00:00:00" json:"got_at"`
	Data      string    `gorm:"column:data" json:"data"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName Room's table name
func (*Room) TableName() string {
	return TableNameRoom
}
