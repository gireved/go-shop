package config

import (
	"github.com/spf13/viper"
	"os"
	"time"
)

var Config *Conf

type Mysql struct {
	// 基础配置
	Dialect  string `yaml:"dialect"`
	DBHost   string `yaml:"dbHost"`
	DBPort   string `yaml:"dbPort"`
	DBName   string `yaml:"dbName"`
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`

	// 连接池配置
	MaxOpenConns    int           `yaml:"maxOpenConns"`    // 最大打开连接数
	MaxIdleConns    int           `yaml:"maxIdleConns"`    // 最大空闲连接数
	ConnMaxLifetime time.Duration `yaml:"connMaxLifetime"` // 连接最大存活时间
	ConnMaxIdleTime time.Duration `yaml:"connMaxIdleTime"` // 连接最大空闲时间

	// 读写分离配置
	WriteDSN     string `yaml:"writeDSN"`     // 主库DSN
	ReadReplica1 string `yaml:"readReplica1"` // 读副本1
	ReadReplica2 string `yaml:"readReplica2"` // 读副本2

	//
}

type Conf struct {
	Mysql map[string]*Mysql `yaml:"mysql"`
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	// viper.AddConfigPath(workDir + "/config/locales")
	viper.AddConfigPath(workDir)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&Config)
	if err != nil {
		panic(err)
	}
}
