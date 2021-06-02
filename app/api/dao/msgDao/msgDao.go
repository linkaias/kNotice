package msgDao

import (
	"gorm.io/gorm"
	"kNotice/app/common/model"
	"kNotice/utils/db"
)

type msgDao struct {
	Db *gorm.DB
}

//NewMsgDao -
func NewMsgDao() *msgDao {
	return &msgDao{
		Db: db.NewDB(),
	}
}

//InsertMsg 添加新消息 返回id
func (t *msgDao) InsertMsg(title, info string) (int, error) {
	msg := &model.MsgModel{
		Title: title,
		Info:  info,
	}
	Db := t.Db
	res := Db.Create(msg)
	return int(msg.ID), res.Error
}
