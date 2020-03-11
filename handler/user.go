package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	user "github.com/ingirls/user/proto/user"
)

// User struct
type User struct{}

// Say Say
func (e *User) Say(ctx context.Context, req *user.Request, rsp *user.Response) error {
	log.Info("Received User.Say request")
	rsp.Msg = "Hello " + req.Name
	return nil
}
