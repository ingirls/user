package config

import (
	"github.com/ingirls/common/cache"
	"github.com/ingirls/common/mysql"
	"github.com/micro/go-micro/v2/config"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-plugins/config/source/consul/v2"
)

const (
	// ProjectNamespace ProjectNamespace
	ProjectNamespace = "ingirls"
	// ServiceNamespace ServiceNamespace
	ServiceNamespace = "user"
)

// InitConfig InitConfig
func InitConfig(consulAddr string) {
	consulSource := consul.NewSource(
		consul.WithAddress(consulAddr),
		consul.WithPrefix("/ingirls"),
		consul.StripPrefix(false),
	)

	conf, _ := config.NewConfig()
	err := conf.Load(consulSource)
	if err != nil {
		log.Error("consul config load fail. err=", err)
	}

	log.Info(conf.Map())

	if err := conf.Get(ProjectNamespace).Scan(&Conf); err != nil {
		log.Error("consul config scan fail. err=", err)
	}

	log.Info(Conf)
}

// Conf Conf
var Conf = &Config{
	Mysql: map[string]mysql.Conf{},
	Redis: map[string]cache.Conf{},
}

// Config Config
type Config struct {
	Mysql map[string]mysql.Conf
	Redis map[string]cache.Conf
}

// MysqlConf MysqlConf
type MysqlConf struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Db       string `json:"db"`
	Port     string `json:"port"`
}

// RedisConf RedisConf
type RedisConf struct {
	Host     string `json:"host"`
	Password string `json:"password"`
	Port     string `json:"port"`
}
