package global

import (
	ut "github.com/go-playground/universal-translator"
	"zhaoshop-api/order-web/config"
	"zhaoshop-api/order-web/proto"
)

var (
	Trans              ut.Translator
	ServerConfig       = &config.ServerConfig{}
	GoodsSrvClient     proto.GoodsClient
	OrderSrvClient     proto.OrderClient
	InventorySrvClient proto.InventoryClient
	NacosConfig        = &config.NacosConfig{}
)
