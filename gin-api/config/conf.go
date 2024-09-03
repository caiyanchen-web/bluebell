// 读取配置文件
package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// config 总配置文件
type config struct {
	System system `yaml:"system"`
	Logger logger `yaml:"logger"`
	Mysql  mysql  `yaml:"mysql"`
	Redis  redis  `yaml:"redis"`
	Token  token  `yaml:"token"`
	Upload upload `yaml:"upload"`
}

// 系统配置
type system struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Env  string `yaml:"env"`
}

// logger日志配置
type logger struct {
	Level        string `yaml:"level"`
	Prefix       string `yaml:"prefix"`
	Director     string `yaml:"director"`
	ShowLine     bool   `yaml:"show_Line"`
	LogInConsole bool   `yaml:"log_In_Console"`
}

// mysql配置
type mysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Db       string `yaml:"db"`
	UserName string `yaml:"username"`
	PassWord string `yaml:"password"`
	LogLevel string `yaml:"log_level"`
	Charset  string `yaml:"charset"`
	MaxIdel  int    `yaml:"maxIdel"`
	MaxOpen  int    `yaml:"maxOpen"`
}

// redis配置
type redis struct {
	Address  string `yaml:"address"`
	PassWord string `yaml:"password"`
	Db       int    `yaml:"db"`
}

// token配置
type token struct {
	Headers    string `yaml:"headers"`
	ExpireTime int    `yaml:"expireTime"`
	Secret     string `yaml:"secret"`
	Issuer     string `yaml:"issuer"`
}
type upload struct {
	Path string `yaml:"path"`
}

var Config *config

// init初始化配置
func init() {
	//读取配置文件config.yaml
	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("换成日志输出")
	}
	err = viper.Unmarshal(Config)
}
