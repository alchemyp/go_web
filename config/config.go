package config

import "gopkg.in/ini.v1"

var Conf = new(AppConfig)

//App配置映射
type AppConfig struct {
	*LogConfig   `json:"log" ini:"log"`
	*MysqlConfig `json:"mysql" ini:"mysql"`
	*RedisConfig `json:"redis" ini:"redis"`
}

//Mysql配置映射
type MysqlConfig struct {
	Host         string `json:"host" ini:"host"`
	Port         int    `json:"port" ini:"port"`
	UserName     string `json:"username" ini:"username"`
	DB           string `json:"db" ini:"db"`
	PassWord     string `json:"password" ini:"password"`
	MaxOpenConns int    `json:"max_open_conns" ini:"max_open_conns"`
	MaxIdleConns int    `json:"max_idle_conns" ini:"max_idle_conns"`
}

//Redis配置映射
type RedisConfig struct {
	Host     string `json:"host" ini:"host"`
	Port     string `json:"port" ini:"port"`
	PassWord string `json:"password" ini:"password"`
	DB       string `json:"db" ini:"db"`
}

//日志配置的映射
type LogConfig struct {
	Level      string `json:"level" ini:"level"`
	FileName   string `json:"filename" ini:"filename"`
	MaxSize    int    `json:"max_size" ini:"max_size"`
	MaxAge     int    `json:"max_age" ini:"max_age"`
	MaxBackups int    `json:"max_backups" ini:"max_backups"`
}

//定义加载文件，映射到结构体的函数
func LoadFromFile(cfgConfig string) (err error) {
	return ini.MapTo(Conf, cfgConfig)
}
