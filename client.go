package client

import (
	"sync"
	"time"

	"github.com/LukmanulHakim18/gorooster-client/helpers"
	"github.com/LukmanulHakim18/gorooster-client/implementors"
	"github.com/LukmanulHakim18/gorooster-client/models"

	"github.com/go-redis/redis/v8"
)

type Gorooster interface {
	// GetEvent returning ttl and event data.
	// And if data stil any in redis
	GetEvent(key string, target interface{}) (ttl time.Duration, err error)

	// SetEvent insert your event to redis.
	// And waiting to fire
	SetEvent(key string, expired time.Duration, event models.Event) error

	// UpdateExpiredEvent for update ttl.
	// And rescheduling your event to fire
	UpdateExpiredEvent(key string, expired time.Duration) error

	// Update data event if still exist in redis
	UpdateDataEvent(key string, event models.Event) error

	// Deleted your event if still exist in redis
	// And cancel your event to fire
	DeleteEvent(key string) error
}

var (
	goroosterRedis *implementors.GoroosterRedisImpl
	once           sync.Once
)

func GetRedisClient(clientName, host, pass string, db int) Gorooster {
	if ok := helpers.ValidatorClinetNameAndKey(clientName); !ok {
		panic("client name can not contain ':' ")
	}
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
