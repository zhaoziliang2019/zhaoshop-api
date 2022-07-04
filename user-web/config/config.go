package config

type ServerConfig struct {
	Host       string       `mapstructure:"host" json:"host"`
	Name       string       `mapstructure:"name" json:"name"`
	Port       int          `mapstructure:"port" json:"port"`
	Tags       []string     `mapstructure:"tags" json:"tags"`
	UserSrv    string       `mapstructure:"user_srv_name" json:"user_srv_name"`
	JWTInfo    JWTConfig    `mapstructure:"jwt" json:"jwt"`
	AliSmsInfo AliSmsConfig `mapstructure:"sms" json:"sms"`
	RedisInfo  RedisConfig  `mapstructure:"redis" json:"redis"`
	ConsulInfo ConsulConfig `mapstructure:"consul" json:"consul"`
}

//jwtconfig
type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}

//阿里短信配置
type AliSmsConfig struct {
	ApiKey     string `mapstructure:"key" json:"key"`
	ApiSecrect string `mapstructure:"secrect" json:"secrect"`
	Expire     int    `mapstructure:"expire" json:"expire"` //过期时间
}

//redis配置
type RedisConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
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
