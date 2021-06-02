package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	userDao2 "kNotice/app/api/dao/userDao"
	"kNotice/utils/help"
	"net/http"
	"strconv"
	"strings"
)

func CheckConnect(c *gin.Context) (uid int, err error) {
	str := c.DefaultQuery("u-id", "")
	str = help.FilteredSQLInject(str)
	if str == "" {
		return 0, errors.New("u-id null")
	}
	//截取id和key
	info := strings.Split(str, "-")
	if len(info) < 2 {
		return 0, errors.New("u-id error")
	}
	uidStr := info[0]
	key := info[1]
	intId, _ := strconv.Atoi(uidStr)
	userDao := userDao2.NewUserDao()
	user, err := userDao.GetUserById(intId)
	if err != nil {
		return 0, err
	}
	if user.Key != key {
		return 0, errors.New("key error")
	}
	return intId, nil
}

//ws://127.0.0.1:9090/api/connect?u-id=2-40542166

// Ws 连接客户端
func Ws(c *gin.Context) {
	id, err := CheckConnect(c)
	if err != nil {
		fmt.Println("u-id err", err)
		return
	}
	//升级协议 用户验证
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	//处理conn
	go Process(conn, id)
}
