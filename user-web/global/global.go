package global

import (
	ut "github.com/go-playground/universal-translator"
	"zhaoshop-api/user-web/config"
	"zhaoshop-api/user-web/proto"
)

var (
	Trans         ut.Translator
	ServerConfig  *config.ServerConfig = &config.ServerConfig{}
	UserSrvClient proto.UserClient
	NacosConfig   *config.NacosConfig = &config.NacosConfig{}
)
