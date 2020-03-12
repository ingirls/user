package cache

import (
	"fmt"

	"log"

	"github.com/go-cache/cache"
	"github.com/go-redis/redis/v7"
)

// Conf struct
type Conf struct {
	Password string
	Host     string
	Port     string
}

// CacheDriver CacheDriver
var CacheDriver map[string]*cache.Client

// With With
func With(name ...string) *cache.Client {
	k := "default"
	if len(name) > 0 {
		k = name[0]
	}

	if _, ok := CacheDriver[k]; ok {
		return CacheDriver[k]
	}

	log.Panicf("redis [%s] not exist.", k)
	return nil
}

// InitCache InitCache
func InitCache(confs map[string]Conf) {
	CacheDriver = make(map[string]*cache.Client)

	for name, conf := range confs {
		redisConf := redis.Options{
			Addr:     fmt.Sprintf("%s:%s", conf.Host, conf.Port),
			Password: conf.Password,
		}

		client := cache.NewClient(redisConf)
		_, err := client.RedisClient.Ping().Result()

		if err != nil {
			log.Fatalf("redis [%s] connect fail. err [%s]", name, err)
			continue
		}

		log.Printf("redis [%s] connected.", name)
		CacheDriver[name] = client
	}
}
