package model

import "time"

//UserModel 用户
type UserModel struct {
	ID        uint      `json:"id" gorm:"primaryKey column:id" form:"id"`
	UserName  string    `json:"username"  gorm:"index;column:username" form:"username"`
	LoginNum  int64     `json:"login_num" gorm:"column:login_num"`
	Status    byte      `json:"status" gorm:"index;column:status" form:"status"`
	Key       string    `json:"key" gorm:"column:key" form:"key"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at" `
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at" `
}

// TableName  自定义表名
func (t *UserModel) TableName() string {
	return "k_users"
}
