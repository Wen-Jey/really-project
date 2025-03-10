//第一部分 2025-03-03

//文件配置

package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

//总配置文件

type config struct {
	//启动配置
	Server server `yaml:"server"`
	//数据库服务
	Db db `yaml:"db"`
	//redis配置
	Redis redis `yaml:"redis"`
	//日志操作
	Log log `yaml:"log"`
	//图片配置
	ImageSettings imageSettings `yaml:"imageSetting"`
}

// 项目的启动端口配置  对应config.yaml中的server
type server struct {
	Address string `yaml:"address"`
	Model   string `yaml:"model"`
}

// 数据库配置  第二部分
type db struct {
	Dialects string `yaml:"dialects"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Db       string `yaml:"db"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
	MaxIdle  int    `yaml:"maxIdle"`
	MaxOpen  int    `yaml:"maxOpen"`
}

// redis配置 第三部分
type redis struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
}

// imageSettings图片上传配置第四部分
type imageSettings struct {
	UploadDir string `yaml:"uploadDir"`
	ImageHost string `yaml:"imageHost"`
}

// log日志配置 第五部分
type log struct {
	Path  string `yaml:"path"`
	Name  string `yaml:"name"`
	Model string `yaml:"model"`
}

var Config *config

//配置服务初始化 用来加载config.yaml文件

func init() {
	yamlFile, err := ioutil.ReadFile("./config.yaml")
	//有错误就进行宕机，停止服务
	if err != nil {
		panic(err)
	}

	//绑定值

	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		panic(err)
	}
}
