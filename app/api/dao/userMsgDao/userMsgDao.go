package userMsgDao

import (
	"gorm.io/gorm"
	"kNotice/app/common/model"
	"kNotice/utils/db"
)

type userMsgDao struct {
	Db *gorm.DB
}

func NewUserMsgDao() *userMsgDao {
	return &userMsgDao{
		Db: db.NewDB(),
	}
}

type msgUserReturn struct {
	model.MsgModel
	Time string `json:"time"`
}

//GetHistoryNotice 获取历史通知消息
func (t *userMsgDao) GetHistoryNotice(pageSize, page, userId int) []*msgUserReturn {
	Db := t.Db
	msgs := make([]*msgUserReturn, 0)
	Db = Db.Table("`k_user_msg` as um").Joins("left join `k_msgs` as m on um.msg_id = m.id").Select("m.*").Where("um.user_id = ?", userId)

	//分页操作
	if page > 0 && pageSize > 0 {
		Db = Db.Limit(pageSize).Offset((page - 1) * pageSize)
	}
	_ = Db.Order("um.`id` desc").Find(&msgs)
	for _, msg := range msgs {
		msg.Time = msg.CreatedAt.Format("2006-01-02 15:04")
	}

	return msgs
}

//SaveUserMsg 保存用户通知
func (t *userMsgDao) SaveUserMsg(userId, msgId uint) {
	Db := t.Db
	userMsg := &model.UserMsgModel{
		UserId: userId,
		MsgId:  msgId,
	}
	Db.Create(userMsg)
}
