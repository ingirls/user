package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	"github.com/ingirls/common/api"
	proto "github.com/ingirls/user/proto/user"
	apiProto "github.com/micro/go-micro/v2/api/proto"
)

// Handler struct
type User struct {
	Client proto.Service
}

// Call is a single request handler called via client.Call or the generated client code
func (h *User) Say(ctx context.Context, req *apiProto.Request, rsp *apiProto.Response) error {

	var args proto.Request
	if err := api.Bind(req, &args); err != nil {
		return err
	}

	log.Info("args ", args)

	resp, err := h.Client.Say(ctx, &proto.Request{Name: args.Name})
	if err != nil {
		return err
	}

	log.Info("resp ", resp)

	return api.JSON(rsp, 200, resp)
}
