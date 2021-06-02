package app

import (
	"github.com/gin-gonic/gin"
	"kNotice/app/api"
)

// SetupRouter router
func SetupRouter() *gin.Engine {
	r := gin.New()

	//api
	r = api.RouterApi(r)

	return r
}
