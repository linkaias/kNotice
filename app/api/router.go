package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kNotice/app/api/controller"
	api2 "kNotice/app/api/controller/api"
)

//CORS 跨域问题
func CORS() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 允许 Origin 字段中的域发送请求
		context.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		// 设置预验请求有效期为 86400 秒
		context.Writer.Header().Set("Access-Control-Max-Age", "86400")
		// 设置允许请求的方法
		context.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE, PATCH")
		// 设置允许请求的 Header
		context.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length，Apitoken")
		// 设置拿到除基本字段外的其他字段，如上面的Apitoken, 这里通过引用Access-Control-Expose-Headers，进行配置，效果是一样的。
		context.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Headers")
		// 配置是否可以带认证信息
		context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		// OPTIONS请求返回200
		if context.Request.Method == "OPTIONS" {
			fmt.Println(context.Request.Header)
			context.AbortWithStatus(200)
		} else {
			context.Next()
		}
	}
}

func RouterApi(r *gin.Engine) *gin.Engine {

	r.Use(CORS())

	api := r.Group("api")
	{
		api.GET("connect", controller.Ws)

		//用户相关
		user := api.Group("user")
		{
			//创建用户
			user.GET("register", api2.CreateUser)
		}
		//notice
		notice := api.Group("notice")
		{
			//发送消息
			notice.GET("send", api2.SendNotice)

			//获取历史消息
			notice.GET("getHistoryNotice", api2.GetHistoryNotice)
		}

		api.GET("test", func(c *gin.Context) {

		})

	}

	return r
}
