package global

import (
	ut "github.com/go-playground/universal-translator"
	"zhaoshop-api/goods-web/config"
	"zhaoshop-api/goods-web/proto"
)

var (
	Trans          ut.Translator
	ServerConfig   = &config.ServerConfig{}
	GoodsSrvClient proto.GoodsClient
	NacosConfig    = &config.NacosConfig{}
)
