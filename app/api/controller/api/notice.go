package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kNotice/app/api/controller"
	"kNotice/app/api/dao/userMsgDao"
	"kNotice/app/common/global"
	"kNotice/app/common/message"
	"kNotice/utils/help"
	"net/http"
	"strconv"
	"time"
)

func SendNotice(c *gin.Context) {
	res := message.RequestMsg{}

	title := c.Query("title")
	title = help.FilteredSQLInject(title)
	msg := c.Query("notice")
	msg = help.FilteredSQLInject(msg)
	if msg != "" {
		not := &message.Notice{
			Title: title,
			Info:  msg,
			Time:  time.Now().Format("2006-01-02 15:04"),
		}
		global.NoticeGlobal <- not
		res.Code = http.StatusOK
		res.Msg = "success!"
	} else {
		res.Code = 100
		res.Msg = "notice nil"
	}

	c.JSON(http.StatusOK, res)
}

//GetHistoryNotice 获取历史通知
func GetHistoryNotice(c *gin.Context) {
	c.Set("Access-Control-Allow-Origin", "*") //允许访问所有域

	res := &message.RequestMsg{}
	uid, err := controller.CheckConnect(c)
	if err != nil {
		fmt.Println("u-id err", err)
		res.Code = 100
		res.Msg = fmt.Sprintf("u-id err:%s", err)
	} else {
		pageStr := c.DefaultQuery("page", "1")
		page, _ := strconv.Atoi(pageStr)
		limitStr := c.DefaultQuery("limit", "5")
		limit, _ := strconv.Atoi(limitStr)

		umDao := userMsgDao.NewUserMsgDao()
		msgs := umDao.GetHistoryNotice(limit, page, uid)
		res.Code = http.StatusOK
		res.Msg = "Success"
		res.Res = msgs
	}
	c.JSON(http.StatusOK, res)
}
