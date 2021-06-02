package model

import "time"

//MsgModel 所有通知
type MsgModel struct {
	ID        uint      `json:"id" gorm:"primaryKey column:id" form:"id"`
	Title     string    `json:"title"  gorm:"column:title" form:"title"`
	Info      string    `json:"info"  gorm:"column:info" form:"info"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at" `
}

// TableName  自定义表名
func (t *MsgModel) TableName() string {
	return "k_msgs"
}
