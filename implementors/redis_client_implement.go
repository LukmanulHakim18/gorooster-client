package implementors

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/LukmanulHakim18/gorooster-client/models"

	"github.com/go-redis/redis/v8"
)

type GoroosterRedisImpl struct {
	DB         *redis.Client
	ClientName string
}

func (res GoroosterRedisImpl) GetEvent(key string, target any) (ttl time.Duration, err error) {
	ctx := context.Background()
	fmt.Println(res.generateKeyEvent(key))
	val := res.DB.Get(ctx, res.generateKeyData(key)).Val()
	if err = json.Unmarshal([]byte(val), target); err != nil {
		return
	}
	ttl, err = res.DB.TTL(ctx, res.generateKeyEvent(key)).Result()
	return
}

func (res GoroosterRedisImpl) SetEvent(key string, expired time.Duration, event models.Event) error {
	ctx := context.Background()
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}
	if err := res.DB.Set(ctx, res.generateKeyEvent(key), "event-key", expired).Err(); err != nil {
		return err
	}
	if err := res.DB.Set(ctx, res.generateKeyData(key), string(data), -1).Err(); err != nil {
		return err
	}
	return nil
}

func (res GoroosterRedisImpl) UpdateExpiredEvent(key string, expired time.Duration) error {
	ctx := context.Background()
	err := res.DB.Get(ctx, res.generateKeyData(key)).Err()
	if err != nil {
		return fmt.Errorf("data not found")
	}
	if err := res.DB.Set(ctx, res.generateKeyEvent(key), "event-key", expired).Err(); err != nil {
		return err
	}
	return nil
}

func (res GoroosterRedisImpl) UpdateDataEvent(key string, event models.Event) error {
	ctx := context.Background()
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}
	if err := res.DB.Get(ctx, res.generateKeyData(key)).Err(); err != nil {
		return err
	}
	return res.DB.Set(ctx, res.generateKeyData(key), data, -1).Err()

}

func (res GoroosterRedisImpl) DeleteEvent(key string) error {
	ctx := context.Background()
	if err := res.DB.Del(ctx, res.generateKeyEvent(key)).Err(); err != nil {
		return err
	}
	if err := res.DB.Del(ctx, res.generateKeyData(key)).Err(); err != nil {
		return err
	}
	return nil
}

func (res GoroosterRedisImpl) generateKeyEvent(key string) string {
	keyEvent := fmt.Sprintf("%s:event:%s", res.ClientName, key)
	return keyEvent
}

func (res GoroosterRedisImpl) generateKeyData(key string) string {
	keyData := fmt.Sprintf("%s:data:%s", res.ClientName, key)
	return keyData
}
