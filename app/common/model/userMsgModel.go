package model

import "time"

//UserMsgModel 用户接收的通知
type UserMsgModel struct {
	ID        uint      `json:"id" gorm:"primaryKey column:id" form:"id"`
	UserId    uint      `json:"user_id" gorm:"index;column:user_id" form:"user_id"`
	MsgId     uint      `json:"msg_id" gorm:"index;column:msg_id" form:"msg_id"`
	Status    byte      `json:"status" gorm:"index;column:status" form:"status"` // 0 未推送
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at" `
}

// TableName  自定义表名
func (t *UserMsgModel) TableName() string {
	return "k_user_msg"
}
