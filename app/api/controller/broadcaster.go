package controller

import (
	"fmt"
	msgDao2 "kNotice/app/api/dao/msgDao"
	userDao2 "kNotice/app/api/dao/userDao"
	userMsgDao2 "kNotice/app/api/dao/userMsgDao"
	"kNotice/app/common/global"
	"kNotice/app/common/message"
)

//Broadcaster 取出待推送消息
func Broadcaster() {
	for {
		select {
		case notice := <-global.NoticeGlobal:
			sendNotice(notice)
		}
	}
}

//推送消息
func sendNotice(notice *message.Notice) {
	go saveNoticeLog(notice.Title, notice.Info)

	for _, client := range global.ClientsGlobal {
		client.Conn.WriteJSON(notice)
	}

}

//保存通知记录
func saveNoticeLog(title, info string) {
	msgDao := msgDao2.NewMsgDao()
	msgId, err := msgDao.InsertMsg(title, info)
	if err != nil {
		fmt.Println("save notice err:", err)
		return
	}
	userDao := userDao2.NewUserDao()
	users, err := userDao.GetAllUser()
	if err != nil {
		fmt.Println("get all users err:", err)
		return
	}
	userMsgDao := userMsgDao2.NewUserMsgDao()
	for _, user := range users {
		userMsgDao.SaveUserMsg(user.ID, uint(msgId))
	}
}
