package main

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/server"

	"github.com/ingirls/user/api/handler"
	proto "github.com/ingirls/user/proto/user"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.user"),
		micro.Version("latest"),
	)

	service.Server().Init(
		server.Wait(nil),
	)

	// Initialise service
	service.Init()

	// Register Handler
	service.Server().Handle(
		service.Server().NewHandler(
			&handler.User{
				Client: proto.NewService("go.micro.srv.user", service.Client()),
			},
		),
	)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}