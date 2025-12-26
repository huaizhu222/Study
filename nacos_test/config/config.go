package config

type UserSrvConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int32  `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}

type AliSmsConfig struct {
	AccessKeyId     string `mapsturcture:"keyid" json:"keyid"`
	AccessKeySecret string `mapsturcture:"keysecret" json:"keysecret"`
}
type RedisConfig struct {
	Host   string `mapstructure:"host" json:"host"`
	Port   int32  `mapstructure:"port" json:"port"`
	Expire int    `mapstructure:"expire" json:"expire"`
}
type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port string `mapstructure:"port" json:"port"`
}
type ServerConfig struct {
	Name          string `mapstructure:"name" json:"name"`
	Port          int32  `mapstructure:"port" json:"port"`
	UserSrvConfig `mapstructure:"user_srv" json:"user_srv"`
	JWTConfig     `mapstructure:"jwt" json:"jwt"`
	AliSmsConfig  `mapstructure:"sms" json:"sms"`
	RedisConfig   `mapstructure:"redis" json:"redis"`
	ConsulConfig  `mapstructure:"consul" json:"consul"`
}
