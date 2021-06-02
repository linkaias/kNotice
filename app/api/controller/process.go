package controller

import (
	"fmt"
	"github.com/gorilla/websocket"
	userDao2 "kNotice/app/api/dao/userDao"
	"kNotice/app/api/model"
	"kNotice/app/common/global"
)

// Process 处理连接
func Process(conn *websocket.Conn, id int) {

	userDao := userDao2.NewUserDao()
	user, err := userDao.GetUserById(id)
	if err != nil {
		conn.Close()
		return
	}
	clint := &model.Client{
		Conn: conn,
		User: user,
	}

	if _, ok := global.ClientsGlobal[int(user.ID)]; !ok {
		global.ClientsGlobal[int(user.ID)] = clint
	}

	//失去连接时处理的事情
	defer func() {
		//去除在线状态
		if _, ok := global.ClientsGlobal[int(user.ID)]; ok {
			delete(global.ClientsGlobal, int(user.ID))
		}
		//关闭链接
		conn.Close()
	}()

	for {
		//一直读取消息
		_, msgStr, err := conn.ReadMessage()
		if err != nil {
			break
		}
		fmt.Println("ws:", string(msgStr))
	}
}
