package conf

import (
	"fmt"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"maxblog-be-template/internal/core"
	"sync"
)

var cfg *Config
var once sync.Once

func init() {
	once.Do(func() {
		cfg = &Config{}
	})
}

func GetInstanceOfConfig() *Config {
	return cfg
}

type Config struct {
	mux     sync.Mutex
	RunMode string `mapstructure:"run_mode" json:"run_mode"`
	Logger  Logger `mapstructure:"logger" json:"logger"`
	App     App    `mapstructure:"app" json:"app"`
	Server  Server `mapstructure:"server" json:"server"`
	DB      DB     `mapstructure:"db" json:"db"`
	Mysql   Mysql  `mapstructure:"mysql" json:"mysql"`
}

type Logger struct {
	Color bool `mapstructure:"color" json:"color"`
}

type App struct {
	AppName string `mapstructure:"app_name" json:"app_name"`
	Version string `mapstructure:"version" json:"version"`
}

type Server struct {
	Host            string `mapstructure:"host" json:"host"`
	Port            int    `mapstructure:"port" json:"port"`
	ShutdownTimeout int    `mapstructure:"shutdown_timeout" json:"shutdown_timeout"`
}

type DB struct {
	Type         string `mapstructure:"type" json:"type"`
	Debug        bool   `mapstructure:"debug" json:"debug"`
	DSN          string `mapstructure:"dsn" json:"dsn"`
	MaxLifetime  int    `mapstructure:"max_life_time" json:"max_life_time"`
	MaxOpenConns int    `mapstructure:"max_open_conns" json:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns" json:"max_idle_conns"`
}

type Mysql struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     string `mapstructure:"port" json:"port"`
	UserName string `mapstructure:"user_name" json:"user_name"`
	Password string `mapstructure:"password" json:"password"`
	DBName   string `mapstructure:"db_name" json:"db_name"`
	Params   string `mapstructure:"params" json:"params"`
}

func (mysql *Mysql) DSN() string {
	if mysql.Password == "" {
		mysql.Password = "123456"
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		mysql.UserName, mysql.Password, mysql.Host, mysql.Port, mysql.DBName, mysql.Params)
}

func (cfg *Config) Load(configFile string) {
	configPath := configFile
	v := viper.New()
	v.SetConfigFile(configPath)
	err := v.ReadInConfig()
	if err != nil {
		logger.WithFields(logger.Fields{
			"失败方法": core.GetFuncName(),
		}).Panic(core.FormatError(900, err).Error())
	}
	err = v.Unmarshal(cfg)
	if err != nil {
		logger.WithFields(logger.Fields{
			"失败方法": core.GetFuncName(),
		}).Panic(core.FormatError(901, err).Error())
	}
}
