package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	userDao2 "kNotice/app/api/dao/userDao"
	"kNotice/app/common/message"
	"kNotice/app/common/model"
	"kNotice/utils/help"
	"net/http"
)

func CreateUser(c *gin.Context) {
	res := message.RequestMsg{}
	userName := c.Query("username")
	if userName == "" {
		res.Code = 100
		res.Msg = "user name null"
	} else {
		key := help.CreateValidateCode(8)
		user := &model.UserModel{
			UserName: userName,
			Status:   1,
			Key:      key,
		}
		userDao := userDao2.NewUserDao()
		if user2, err := userDao.CreateUser(user); err != nil {
			res.Code = 100
			res.Msg = "save table err"
		} else {
			result := make(map[string]string)
			result["id"] = fmt.Sprintf("%d", user2.ID)
			result["username"] = user2.UserName
			result["u-id"] = fmt.Sprintf("%d-%s", user2.ID, user2.Key)
			res.Res = result
		}
	}
	c.JSON(http.StatusOK, res)
}
