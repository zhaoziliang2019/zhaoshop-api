package global

import (
	ut "github.com/go-playground/universal-translator"
	"zhaoshop-api/userop-web/config"
	"zhaoshop-api/userop-web/proto"
)

var (
	Trans ut.Translator

	ServerConfig = &config.ServerConfig{}

	NacosConfig = &config.NacosConfig{}

	GoodsSrvClient proto.GoodsClient

	MessageClient proto.MessageClient
	AddressClient proto.AddressClient
	UserFavClient proto.UserFavClient
)
