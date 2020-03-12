package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	"github.com/ingirls/user/model"
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

// Create Create
func (e *User) Create(ctx context.Context, info *user.User, rsp *user.User) error {
	user, err := model.MysqlModel.Create(info)
	if err != nil {
		log.Errorf("MysqlModel.Create fail. info=%v err=%v", info, err)
		return err
	}

	user.Password = ""
	user.Salt = ""

	*rsp = *user
	return nil
}

// List List
func (e *User) List(ctx context.Context, req *user.Request, rsp *user.Response) error {
	log.Info("Received User.List request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Delete Delete
func (e *User) Delete(ctx context.Context, req *user.User, rsp *user.Null) error {
	user, err := model.MysqlModel.Delete(req)
	if err != nil {
		log.Errorf("MysqlModel.Delete fail. req=%v err=%v", req, err)
		return err
	}

	user.Password = ""
	user.Salt = ""

	return nil
}

// Update Update
func (e *User) Update(ctx context.Context, req *user.User, rsp *user.User) error {
	user, err := model.MysqlModel.Update(req)
	if err != nil {
		log.Errorf("MysqlModel.Update fail. req=%v err=%v", req, err)
		return err
	}

	user.Password = ""
	user.Salt = ""

	*rsp = *user
	return nil
}

// InfoByID InfoByID
func (e *User) InfoByID(ctx context.Context, req *user.ID, rsp *user.User) (err error) {
	user, err := model.CacheModel.InfoByID(req)
	if err != nil {
		log.Errorf("CacheModel.InfoByID fail. req=%v err=%v", req, err)
		return err
	}

	if user == nil || user.Id == 0 {
		return nil
	}

	*rsp = *user
	return nil
}

// InfoByUsername InfoByUsername
func (e *User) InfoByUsername(ctx context.Context, req *user.Username, rsp *user.User) error {
	user, err := model.CacheModel.InfoByUsername(req)
	if err != nil {
		log.Errorf("CacheModel.InfoByUsername fail. req=%v err=%v", req, err)
		return err
	}

	if user == nil || user.Id == 0 {
		return nil
	}

	*rsp = *user
	return nil
}

// InfoByMobile InfoByMobile
func (e *User) InfoByMobile(ctx context.Context, req *user.Mobile, rsp *user.User) error {
	user, err := model.CacheModel.InfoByMobile(req)
	if err != nil {
		log.Errorf("CacheModel.InfoByMobile fail. req=%v err=%v", req, err)
		return err
	}

	if user == nil || user.Id == 0 {
		return nil
	}

	*rsp = *user
	return nil
}

// InfoByEmail InfoByEmail
func (e *User) InfoByEmail(ctx context.Context, req *user.Email, rsp *user.User) error {
	user, err := model.CacheModel.InfoByEmail(req)
	if err != nil {
		log.Errorf("CacheModel.InfoByEmail fail. req=%v err=%v", req, err)
		return err
	}

	if user == nil || user.Id == 0 {
		return nil
	}

	*rsp = *user
	return nil
}

// CheckPassword CheckPassword
func (e *User) CheckPassword(ctx context.Context, req *user.IDAndPassword, rsp *user.User) error {
	user, err := model.MysqlModel.InfoByID(&user.ID{Id: req.Id})
	if err != nil {
		log.Errorf("MysqlModel.InfoByID fail. req=%v err=%v", req, err)
		return err
	}

	if user == nil || user.Id == 0 {
		return nil
	}

	if !model.PasswordVerify(req.Password, user.Salt, user.Password) {
		log.Info("model.PasswordVerify fail.")
		return nil
	}

	*rsp = *user
	return nil
}

// UpdatePassword UpdatePassword
func (e *User) UpdatePassword(ctx context.Context, req *user.IDAndPassword, rsp *user.User) error {
	// user, err := model.MysqlModel.InfoByID(&user.ID{Id: req.Id})
	// if err != nil {
	// 	log.Errorf("MysqlModel.UpdatePassword fail. req=%v err=%v", req, err)
	// 	return err
	// }

	// if user == nil || user.Id == 0 {
	// 	return nil
	// }

	// if !model.PasswordVerify(req.Password, user.Salt, user.Password) {
	// 	log.Info("model.PasswordVerify fail.")
	// 	return nil
	// }

	// *rsp = *user
	return nil
}
