package config

import (
	"github.com/ingirls/common/cache"
	"github.com/ingirls/common/log"
	"github.com/ingirls/common/mysql"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source"

	"github.com/micro/go-plugins/config/source/consul/v2"
)

const (
	// ProjectNamespace ProjectNamespace
	ProjectNamespace = "ingirls"
	// ServiceNamespace ServiceNamespace
	ServiceNamespace = "user-srv"
	// APINamespace APINamespace
	APINamespace = "user-api"
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

	if err := conf.Get(ProjectNamespace).Scan(&Conf); err != nil {
		log.Error("consul config scan fail. err=", err)
	}

	log.Info("conf load success.")
	go confWatcher(conf, consulSource)
}

func confWatcher(conf config.Config, consulSource source.Source) {
	// 开始侦听变动事件
	watcher, err := conf.Watch()
	if err != nil {
		log.Fatal(err)
	}

	log.Info("watching conf")

	for {
		_, err := watcher.Next()
		if err != nil {
			log.Fatal(err)
		}

		log.Infof("conf is changing...")

		err = conf.Load(consulSource)
		if err != nil {
			log.Error("consul config load fail. err=", err)
		}

		if err := conf.Get(ProjectNamespace).Scan(&Conf); err != nil {
			log.Error("consul config scan fail. err=", err)
		}

		log.Info("new conf load success.")
	}
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
