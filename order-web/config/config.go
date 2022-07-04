package config

type ServerConfig struct {
	Name         string       `mapstructure:"name" json:"name"`
	Host         string       `mapstructure:"host" json:"host"`
	Port         int          `mapstructure:"port" json:"port"`
	Tags         []string     `mapstructure:"tags" json:"tags"`
	GoodsSrv     string       `mapstructure:"goods_srv_name" json:"goods_srv_name"`
	OrderSrv     string       `mapstructure:"order_srv_name" json:"order_srv_name"`
	InventorySrv string       `mapstructure:"inventory_srv_name" json:"inventory_srv_name"`
	JWTInfo      JWTConfig    `mapstructure:"jwt" json:"jwt"`
	ConsulInfo   ConsulConfig `mapstructure:"consul" json:"consul"`
	AlipayInfo   AlipayConfig `mapstructure:"alipay" json:"alipay"`
	JaegerInfo   JaegerConfig `mapstructure:"jaeger" json:"jaeger"`
}

//jwtconfig
type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}

//consul配置
type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

//alipay
type AlipayConfig struct {
	AppID        string `mapstructure:"appid" json:"appid"`
	PrivateKey   string `mapstructure:"private_key" json:"private_key"`
	AliPublicKey string `mapstructure:"ali_public_key" json:"ali_public_key"`
	NotifyUrl    string `mapstructure:"notify_url" json:"notify_url"`
	ReturnUrl    string `mapstructure:"return_url" json:"return_url"`
}

//nacos配置
type NacosConfig struct {
	Host      string `mapstructure:"host"`
	Port      uint64 `mapstructure:"port"`
	Namespace string `mapstructure:"namespace"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	DataId    string `mapstructure:"dataid"`
	Group     string `mapstructure:"group"`
}

//jaeger
type JaegerConfig struct {
	Host string `mapstructure:"host"`
	Port uint64 `mapstructure:"port"`
	Name string `mapstructure:"name" json:"name"`
}
