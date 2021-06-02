package model

import (
	"github.com/gorilla/websocket"
	"kNotice/app/common/model"
)

type Client struct {
	Conn *websocket.Conn
	User *model.UserModel
}
