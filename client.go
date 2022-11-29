package client

import (
	"sync"
	"time"

	"github.com/LukmanulHakim18/gorooster-client/implementors"
	"github.com/LukmanulHakim18/gorooster-client/models"

	"github.com/go-redis/redis/v8"
)

type Gorooster interface {
	GetEvent(key string, target any) (ttl time.Duration, err error)
	SetEvent(key string, expired time.Duration, event models.Event) error
	UpdateExpiredEvent(key string, expired time.Duration) error
	UpdateDataEvent(key string, event models.Event) error
	DeleteEvent(key string) error
}

var (
	goroosterRedis *implementors.GoroosterRedisImpl
	once           sync.Once
)

func GetRedisClient(clientName, host, pass string, db int) Gorooster {
	once.Do(func() {
		if goroosterRedis == nil {
			redisDB := redis.NewClient(&redis.Options{
				Addr:     host,
				DB:       db,
				Password: pass,
			})
			goroosterRedis = &implementors.GoroosterRedisImpl{
				DB:         redisDB,
				ClientName: clientName,
			}
		}
	})
	return goroosterRedis
}
