package main

import (
	"fmt"
	"kNotice/app"
	"kNotice/app/api/controller"
	"kNotice/app/common/global"
)

func init() {
	global.GblInit()
	go controller.Broadcaster()
}

func main() {

	err := app.SetupRouter().Run(":9090")
	if err != nil {
		fmt.Println(err)
		return
	}

}
