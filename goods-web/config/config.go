package config

type ServerConfig struct {
	Name       string       `mapstructure:"name" json:"name"`
	Host       string       `mapstructure:"host" json:"host"`
	Port       int          `mapstructure:"port" json:"port"`
	Tags       []string     `mapstructure:"tags" json:"tags"`
	GoodsInfo  string       `mapstructure:"goods_srv_name" json:"goods_srv_name"`
	JWTInfo    JWTConfig    `mapstructure:"jwt" json:"jwt"`
	ConsulInfo ConsulConfig `mapstructure:"consul" json:"consul"`
	JaegerInfo JaegerConfig `mapstructure:"jaeger" json:"jaeger"`
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
