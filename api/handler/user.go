package handler

import (
	"context"

	"github.com/ingirls/common/log"

	proto "github.com/ingirls/user/api/proto/user"
	user "github.com/ingirls/user/proto/user"
)

// User struct
type User struct {
	Client user.Service
}

// Say is a single request handler called via client.Call or the generated client code
func (h *User) Say(ctx context.Context, req *proto.Request, rsp *proto.Response) error {

	log.Info(req)

	resp, err := h.Client.Say(ctx, &user.Request{Name: req.Name})
	if err != nil {
		return err
	}

	log.Info("resp ", resp)
	rsp.Msg = resp.Msg
	return nil
}
