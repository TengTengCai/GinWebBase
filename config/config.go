package config

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	)

var (
	conf     *Config
)

// 配置文件结构体
type Config struct {
	DataBase DataBase    `yaml:"database"`
	Server	 Server		`yaml:"server"`
	Log		Log			`yaml:"log"`
	// 可添加其他结构体,作为配置解析
}

type DataBase struct {
	Type       string `yaml:"type"`
	ConnectStr string `yaml:"connectStr"`
	Init       bool   `yaml:"init"`
}

type Server struct {
	Addr string `yaml:"addr"`
	Port string `yaml:"port"`
}

type Log struct {
	LogPath string `yaml:"logPath"`
	LogName string `yaml:"logName"`
}


//func GetConfig() *Config {
//	return conf
//}

// 读取配置文件, 解析到结构体中
func Read() *Config {
	// Read config file
	buf, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	// Unmarshal yml
	conf = &Config{}
	err = yaml.Unmarshal(buf, conf)
	if err != nil {
		panic(err)
	}
	return conf
}