package client

import (
	"net/http"
	"sync"
	"time"

	"git.bluebird.id/mybb/gorooster-client/v2/helpers"
	"git.bluebird.id/mybb/gorooster-client/v2/implementors"
	"git.bluebird.id/mybb/gorooster-client/v2/models"
)

type Gorooster interface {
	// GetEvent returning ttl and event data.
	// And if data stil any in redis
	GetEvent(key string, target interface{}) (ttl time.Duration, err error)

	// SetEvent insert your event to redis.
	// And waiting to fire
	SetEvent(key string, eventReleaseIn time.Duration, event models.Event) error
	// SetEvent insert your event to redis.
	// And waiting to fire
	SetEventAt(key string, eventReleaseIn time.Time, event models.Event) error

	// UpdateReleaseEvent for update ttl.
	// And rescheduling your event to fire
	UpdateReleaseEvent(key string, eventReleaseIn time.Duration) error
	// UpdateReleaseEvent for update ttl.
	// And rescheduling your event to fire
	UpdateReleaseEventAt(key string, eventReleaseIn time.Time) error

	// Update data event if still exist in redis
	UpdateDataEvent(key string, event models.Event) error

	// Deleted your event if still exist in redis
	// And cancel your event to fire
	DeleteEvent(key string) error
}

var (
	goroosterAPI *implementors.GoroosterAPIImpl
	onceAPI      sync.Once
)

func GetAPIClient(clientName, baseUrl, apiKey string) Gorooster {
	if ok := helpers.ValidatorClinetNameAndKey(clientName); !ok {
		panic("client name can not contain ':' ")
	}
	onceAPI.Do(func() {
		if goroosterAPI == nil {
			c := http.Client{
				Timeout: time.Duration(2) * time.Second,
			}

			goroosterAPI = &implementors.GoroosterAPIImpl{
				ClientName: clientName,
				Client:     &c,
				BaseUrl:    baseUrl,
			}
		}
	})
	return goroosterAPI
}
