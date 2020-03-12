package model

import (
	"encoding/json"

	"github.com/go-redis/redis/v7"
	"github.com/ingirls/common/cache"
	proto "github.com/ingirls/user/proto/user"
	log "github.com/micro/go-micro/v2/logger"
)

// Cache struct
type Cache struct{}

// CacheModel CacheModel
var CacheModel = &Cache{}

// Update Update
func (m *Cache) Update(info *proto.User) (user *proto.User, err error) {
	user, err = MysqlModel.Update(info)
	if err != nil {
		log.Errorf("MysqlModel.Update fail. info=%v err=%v", info, err)
		return nil, err
	}

	if user == nil || user.Id == 0 {
		return nil, nil
	}

	// 更新缓存 TODO

	return user, nil
}

// Delete Delete
func (m *Cache) Delete(info *proto.User) (user *proto.User, err error) {
	user, err = MysqlModel.Delete(info)
	if err != nil {
		log.Errorf("MysqlModel.Delete fail. info=%v err=%v", info, err)
		return nil, err
	}

	if user == nil || user.Id == 0 {
		return nil, nil
	}

	// 更新缓存 TODO

	return user, nil
}

// InfoByID InfoByID
func (m *Cache) InfoByID(ID *proto.ID) (user *proto.User, err error) {
	res, err := cache.With().RedisClient.Get(GetUserIDKey(ID.Id)).Result()
	if err != nil {
		if err != redis.Nil {
			log.Errorf("RedisClient.Get fail. ID=%v err=%v", ID, err)
			return nil, err
		}
	}

	// 有缓存
	if res != "" {
		if err := json.Unmarshal([]byte(res), &user); err != nil {
			log.Errorf("json.Unmarshal fail. ID=%v err=%v", ID, err)
			return nil, err
		}

		log.Infof("InfoByID from cache. ID=%v", ID)
		return user, nil
	}

	// 从数据库查询
	user, err = MysqlModel.InfoByID(ID)
	if err != nil {
		log.Errorf("MysqlModel.InfoByID fail. ID=%v err=%v", ID, err)
		return nil, err
	}

	if user == nil || user.Id == 0 {
		return nil, nil
	}

	// 写入缓存
	user = ClearUserPassword(user)
	err = cache.With().Tag(GetUserAllTag(), GetUserIDTag(ID.Id)).Put(GetUserIDKey(ID.Id), user, 0)
	if err != nil {
		log.Errorf("Tag.Put fail. ID=%v err=%v", ID, err)
		return nil, err
	}

	log.Infof("InfoByID from db. ID=%v", ID)
	return user, nil
}

// InfoByUsername InfoByUsername
func (m *Cache) InfoByUsername(Username *proto.Username) (user *proto.User, err error) {
	res, err := cache.With().RedisClient.Get(GetUserUsernameKey(Username.Username)).Result()
	if err != nil {
		if err != redis.Nil {
			log.Errorf("RedisClient.Get fail. Username=%v err=%v", Username, err)
			return nil, err
		}
	}

	// 有缓存
	if res != "" {
		if err := json.Unmarshal([]byte(res), &user); err != nil {
			log.Errorf("json.Unmarshal fail. Username=%v err=%v", Username, err)
			return nil, err
		}
		log.Infof("InfoByUsername from cache. Username=%v", Username)
		return user, nil
	}

	// 从数据库查询
	user, err = MysqlModel.InfoByUsername(Username)
	if err != nil {
		log.Errorf("MysqlModel.InfoByUsername fail. Username=%v err=%v", Username, err)
		return nil, err
	}

	if user == nil || user.Id == 0 {
		return nil, nil
	}

	// 写入缓存
	user = ClearUserPassword(user)
	err = cache.With().Tag(GetUserAllTag(), GetUserUsernameTag(Username.Username)).Put(GetUserUsernameKey(Username.Username), user, 0)
	if err != nil {
		log.Errorf("Tag.Put fail. Username=%v err=%v", Username, err)
		return nil, err
	}

	log.Infof("InfoByUsername from db. Username=%v", Username)
	return user, nil
}

// InfoByMobile InfoByMobile
func (m *Cache) InfoByMobile(Mobile *proto.Mobile) (user *proto.User, err error) {
	res, err := cache.With().RedisClient.Get(GetUserMobileKey(Mobile.Mobile)).Result()
	if err != nil {
		if err != redis.Nil {
			log.Errorf("RedisClient.Get fail. Mobile=%v err=%v", Mobile, err)
			return nil, err
		}
	}

	// 有缓存
	if res != "" {
		if err := json.Unmarshal([]byte(res), &user); err != nil {
			log.Errorf("json.Unmarshal fail. Mobile=%v err=%v", Mobile, err)
			return nil, err
		}

		log.Infof("InfoByMobile from cache. Mobile=%v", Mobile)
		return user, nil
	}

	// 从数据库查询
	user, err = MysqlModel.InfoByMobile(Mobile)
	if err != nil {
		log.Errorf("MysqlModel.InfoByMobile fail. Mobile=%v err=%v", Mobile, err)
		return nil, err
	}

	if user == nil || user.Id == 0 {
		return nil, nil
	}

	// 写入缓存
	user = ClearUserPassword(user)
	err = cache.With().Tag(GetUserAllTag(), GetUserMobileTag(Mobile.Mobile)).Put(GetUserMobileKey(Mobile.Mobile), user, 0)
	if err != nil {
		log.Errorf("Tag.Put fail. Mobile=%v err=%v", Mobile, err)
		return nil, err
	}

	log.Infof("InfoByMobile from db. Mobile=%v", Mobile)
	return user, nil
}

// InfoByEmail InfoByEmail
func (m *Cache) InfoByEmail(Email *proto.Email) (user *proto.User, err error) {
	res, err := cache.With().RedisClient.Get(GetUserEmailKey(Email.Email)).Result()
	if err != nil {
		if err != redis.Nil {
			log.Errorf("RedisClient.Get fail. Email=%v err=%v", Email, err)
			return nil, err
		}
	}
	// 有缓存
	if res != "" {
		if err := json.Unmarshal([]byte(res), &user); err != nil {
			log.Errorf("json.Unmarshal fail. Email=%v err=%v", Email, err)
			return nil, err
		}

		log.Infof("InfoByEmail from cache. Email=%v", Email)
		return user, nil
	}

	// 从数据库查询
	user, err = MysqlModel.InfoByEmail(Email)
	if err != nil {
		log.Errorf("MysqlModel.InfoByEmail fail. Email=%v err=%v", Email, err)
		return nil, err
	}

	if user == nil || user.Id == 0 {
		return nil, nil
	}

	// 写入缓存
	user = ClearUserPassword(user)
	err = cache.With().Tag(GetUserAllTag(), GetUserEmailTag(Email.Email)).Put(GetUserEmailKey(Email.Email), user, 0)
	if err != nil {
		log.Errorf("Tag.Put fail. Email=%v err=%v", Email, err)
		return nil, err
	}

	log.Infof("InfoByEmail from db. Email=%v", Email)
	return user, nil
}
