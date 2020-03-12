package main

import (
	"context"
	"net/http"
	"time"

	"github.com/ingirls/common/cache"
	"github.com/ingirls/common/log"
	"github.com/ingirls/common/mysql"
	"github.com/ingirls/user/config"
	"github.com/ingirls/user/handler"
	"github.com/ingirls/user/model"
	"github.com/ingirls/user/subscriber"

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"

	// log "github.com/micro/go-micro/v2/logger"

	"github.com/micro/go-micro/v2/server"
	"github.com/micro/go-plugins/wrapper/monitoring/prometheus/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	user "github.com/ingirls/user/proto/user"
)

func main() {
	log.New("./logs/" + config.ServiceNamespace + ".promtail.log")

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
		micro.Action(action),
		micro.WrapHandler(logWrapper, prometheus.NewHandlerWrapper()),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
		micro.Flags(
			&cli.StringFlag{
				Name:    "consul_addr",
				Usage:   "consul addr",
				EnvVars: []string{"CONSUL_ADDR"},
			},
			&cli.StringFlag{
				Name:    "prometheus_addr",
				Usage:   "prometheus addr",
				EnvVars: []string{"PROMETHEUS_ADDR"},
			},
		),
	)

	service.Server().Init(
		server.Wait(nil),
	)

	// Initialise service
	service.Init()

	// Register Handler
	user.RegisterServiceHandler(service.Server(), new(handler.User))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.user", service.Server(), new(subscriber.User))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func action(c *cli.Context) error {
	// 获取配置
	config.InitConfig(c.String("consul_addr"))

	// 实例化缓存 数据库
	mysql.InitMysql(config.Conf.Mysql)
	cache.InitCache(config.Conf.Redis)

	// 初始化 数据表
	model.AutoMigrate()

	//prometheusBoot
	if len(c.String("prometheus_addr")) > 0 {
		prometheusBoot(c.String("prometheus_addr"))
	}
	return nil
}

// logWrapper is a handler wrapper
func logWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		log.Infof("[wrapper] srv request=%v", req.Endpoint())
		return fn(ctx, req, rsp)
	}
}

func prometheusBoot(addr string) error {
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		err := http.ListenAndServe(addr, nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()
	return nil
}
