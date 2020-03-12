package model

import (
	"time"

	"github.com/jinzhu/gorm"
	log "github.com/micro/go-micro/v2/logger"

	model "github.com/ingirls/common/mysql"
	proto "github.com/ingirls/user/proto/user"

	// mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Mysql struct
type Mysql struct{}

// MysqlModel MysqlModel
var MysqlModel = &Mysql{}

// Create Create
func (m *Mysql) Create(info *proto.User) (user *proto.User, err error) {
	info.CreatedAt = time.Now().Unix()
	info.UpdatedAt = time.Now().Unix()

	if len(info.Password) > 0 {
		info.Password, info.Salt, err = PasswordHash(info.Password)
		if err != nil {
			return nil, err
		}
	}

	info.Status = proto.STATUS_NORMAL

	err = model.With().Create(info).Error

	if err != nil {
		log.Errorf("Mysql.Create fail. info=%v err=%v", info, err)
		return nil, err
	}

	return info, nil
}

// Update Update except password & salt
func (m *Mysql) Update(info *proto.User) (user *proto.User, err error) {
	info.UpdatedAt = time.Now().Unix()

	err = model.With().Omit("password", "salt").Update(info).Error

	if err != nil {
		log.Errorf("Mysql.Update fail. info=%v err=%v", info, err)
		return nil, err
	}

	return info, nil
}

// Delete Delete just update status
func (m *Mysql) Delete(info *proto.User) (user *proto.User, err error) {
	info.UpdatedAt = time.Now().Unix()
	info.Status = proto.STATUS_DELETED

	err = model.With().Omit("password", "salt").Update(info).Error

	if err != nil {
		log.Errorf("Mysql.Delete fail. info=%v err=%v", info, err)
		return nil, err
	}

	return info, nil
}

// InfoByID InfoByID
func (m *Mysql) InfoByID(ID *proto.ID) (user *proto.User, err error) {
	user = &proto.User{}
	err = model.With().Where("id = ?", ID.Id).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Infof("Mysql.InfoByID record not found. ID=%v err=%v", ID, err)
			return nil, nil
		}
		log.Errorf("Mysql.InfoByID fail. ID=%v err=%v", ID, err)
		return nil, err
	}
	return user, nil
}

// InfoByUsername InfoByUsername
func (m *Mysql) InfoByUsername(Username *proto.Username) (user *proto.User, err error) {
	user = &proto.User{}
	err = model.With().Where("username = ?", Username.Username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Infof("Mysql.InfoByUsername record not found. Username=%v err=%v", Username, err)
			return nil, nil
		}
		log.Errorf("Mysql.InfoByUsername fail. Username=%v err=%v", Username, err)
		return nil, err
	}
	return user, nil
}

// InfoByMobile InfoByMobile
func (m *Mysql) InfoByMobile(Mobile *proto.Mobile) (user *proto.User, err error) {
	user = &proto.User{}
	err = model.With().Where("mobile = ?", Mobile.Mobile).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Infof("Mysql.InfoByMobile record not found. Mobile=%v err=%v", Mobile, err)
			return nil, nil
		}
		log.Errorf("Mysql.InfoByMobile fail. Mobile=%v err=%v", Mobile, err)
		return nil, err
	}
	return user, nil
}

// InfoByEmail InfoByEmail
func (m *Mysql) InfoByEmail(Email *proto.Email) (user *proto.User, err error) {
	user = &proto.User{}
	err = model.With().Where("email = ?", Email.Email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Infof("Mysql.InfoByEmail record not found. Email=%v err=%v", Email, err)
			return nil, nil
		}
		log.Errorf("Mysql.InfoByEmail fail. Email=%v err=%v", Email, err)
		return nil, err
	}
	return user, nil
}
