package implementors

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/LukmanulHakim18/gorooster-client/mybb/gorooster-client/v2/helpers"
	"github.com/LukmanulHakim18/gorooster-client/mybb/gorooster-client/v2/models"

	"github.com/go-redis/redis/v8"
)

type GoroosterRedisImpl struct {
	DB         *redis.Client
	ClientName string
}

func (res GoroosterRedisImpl) GetEvent(key string, target interface{}) (ttl time.Duration, err error) {
	if ok := helpers.ValidatorClinetNameAndKey(key); !ok {
		err = fmt.Errorf("key can not contain ':'")
		return
	}
	ctx := context.Background()
	err = res.DB.Get(ctx, helpers.GenerateKeyEvent(res.ClientName, key)).Err()
	if err != nil {
		return 0, fmt.Errorf("data not found")
	}

	val := res.DB.Get(ctx, helpers.GenerateKeyData(res.ClientName, key)).Val()

	if err = json.Unmarshal([]byte(val), target); err != nil {
		return
	}
	ttl, err = res.DB.TTL(ctx, helpers.GenerateKeyEvent(res.ClientName, key)).Result()
	return
}

func (res GoroosterRedisImpl) SetEvent(key string, eventReleaseIn time.Duration, event models.Event) error {
	if ok := helpers.ValidatorClinetNameAndKey(key); !ok {
		return fmt.Errorf("key can not contain ':'")
	}
	ctx := context.Background()
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}
	if err := res.DB.Set(ctx, helpers.GenerateKeyEvent(res.ClientName, key), "event-key", eventReleaseIn).Err(); err != nil {
		return err
	}
	if err := res.DB.Set(ctx, helpers.GenerateKeyData(res.ClientName, key), string(data), -1).Err(); err != nil {
		return err
	}
	return nil
}

func (res GoroosterRedisImpl) UpdateReleaseEvent(key string, eventReleaseIn time.Duration) error {
	if ok := helpers.ValidatorClinetNameAndKey(key); !ok {
		return fmt.Errorf("key can not contain ':'")
	}
	ctx := context.Background()
	err := res.DB.Get(ctx, helpers.GenerateKeyEvent(res.ClientName, key)).Err()
	if err != nil {
		return fmt.Errorf("data not found")
	}
	if err := res.DB.Set(ctx, helpers.GenerateKeyEvent(res.ClientName, key), "event-key", eventReleaseIn).Err(); err != nil {
		return err
	}
	return nil
}

func (res GoroosterRedisImpl) UpdateDataEvent(key string, event models.Event) error {
	if ok := helpers.ValidatorClinetNameAndKey(key); !ok {
		return fmt.Errorf("key can not contain ':'")
	}
	ctx := context.Background()
	err := res.DB.Get(ctx, helpers.GenerateKeyEvent(res.ClientName, key)).Err()
	if err != nil {
		return fmt.Errorf("data not found")
	}
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}
	if err := res.DB.Get(ctx, helpers.GenerateKeyData(res.ClientName, key)).Err(); err != nil {
		return err
	}
	return res.DB.Set(ctx, helpers.GenerateKeyData(res.ClientName, key), data, -1).Err()

}

func (res GoroosterRedisImpl) DeleteEvent(key string) error {
	if ok := helpers.ValidatorClinetNameAndKey(key); !ok {
		return fmt.Errorf("key can not contain ':'")
	}
	ctx := context.Background()
	err := res.DB.Get(ctx, helpers.GenerateKeyEvent(res.ClientName, key)).Err()
	if err != nil {
		return fmt.Errorf("data not found")
	}
	if err := res.DB.Del(ctx, helpers.GenerateKeyEvent(res.ClientName, key)).Err(); err != nil {
		return err
	}
	if err := res.DB.Del(ctx, helpers.GenerateKeyData(res.ClientName, key)).Err(); err != nil {
		return err
	}
	return nil
}
