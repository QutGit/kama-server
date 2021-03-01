package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// DBConfig 数据库链接配置信息
type DBConfig struct {
	Host     string `mapstruct:"host"`
	User     string `mapstruct:"user"`
	Password string `mapstruct:"password"`
	Database string `mapstruct:"database"`
	Port     uint32 `mapstruct:"port"`
}

// Config 声明配置信息所需要的字段
type Config struct {
	Name     string            `mapstruct:"name"`
	Port     uint32            `mapstruct:"port"`
	Database DBConfig          `mapstruct:"database"`
	Service  map[string]string `mapstruct:"service"`
}

// C 配置信息
var C Config

func init() {
	viper.BindEnv("GO_ENV")
	viper.BindEnv("CONFIG_PATH")
	name := GoEnv()
	viper.SetConfigName(name)
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	parentDir := pwd
	for {
		viper.AddConfigPath(filepath.Join(parentDir, "config"))
		log.Println("扫描config路径:", filepath.Join(parentDir, "config"))
		if parentDir == "/" {
			break
		}
		parentDir, _ = filepath.Split(parentDir)
		parentDir = filepath.Dir(parentDir)
	}
	viper.AddConfigPath("./config")
	viper.SetConfigType("yaml")
	err = viper.ReadInConfig()
	if err != nil {
		log.Printf("加载配置信息错误, %s", err)
	}
	viper.Unmarshal(&C)
}

// GoEnv golang环境
func GoEnv() string {
	env := viper.GetString("GO_ENV")
	var envs []string = []string{
		"production",  // 生产
		"development", // 开发
		"stage",       // 预发
		"local",       // 本地，本地调试时，设置环境变量 GO_ENV=local
	}
	for _, v := range envs {
		if v == env {
			return v
		}
	}
	return "production"
}
