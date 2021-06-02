package global

import (
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"kNotice/app/api/model"
	"kNotice/app/common/message"
	"log"
	"os"
	"strings"
)

//ViperGlobal viperr置文件
var ViperGlobal *viper.Viper

//ClientsGlobal 在线用户列表
var ClientsGlobal map[int]*model.Client

//NoticeGlobal 待通知消息管道
var NoticeGlobal chan *message.Notice

//GblInit 初始化
func GblInit() {
	ClientsGlobal = make(map[int]*model.Client)
	NoticeGlobal = make(chan *message.Notice, 10)
	//载入配置
	cfgInit()
}

//载入配置
func cfgInit() {
	path := "config/config.yml"
	viper.SetConfigFile(path)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(fmt.Sprintf("Read config file fail: %s", err.Error()))
	}

	//Replace environment variables
	err = viper.ReadConfig(strings.NewReader(os.ExpandEnv(string(content))))
	if err != nil {
		log.Fatal(fmt.Sprintf("Parse config file fail: %s", err.Error()))
	}

	//global viper
	ViperGlobal = viper.Sub("settings")
}
