# Event Scheduler Client

This repo use for support event scheduler service for set event scheduler.

## Features

- Set event
- Get event
- Update expired event
- Update data event
- Delete event

## Instalation

```go
go get github.com/LukmanulHakim18/gorooster-client
```

## Quickstart

```go
import (
	"github.com/LukmanulHakim18/gorooster-client/models"
	"time"
)

func ExampleClient() {
	dataRaw := map[string]any{
		"location": map[string]any{
			"latitude":  -6.246761122,
			"longitude": 106.8256878,
		},
		"customer_id": "BB12345678",
		"cid":         "021",
	}
	job := models.JobAPI{
		Endpoint: "https://stgapi4.bluebird.id/odrd/trip/newid",
		Method:   models.METHOD_POST,
		Data:     dataRaw,
		Headers: []models.Header{
			models.Header{
				Key:   "Token",
				Value: "mybb-odrd-token",
			},
			models.Header{
				Key:   "Content-Type",
				Value: "application/json",
			},
		},
	}

	dataEvent := models.Event{
		Name:    "cancel_order",
		Id:      "101ec8dc-8de2-448c-b64c-6f0bc49cabff",
		Type:    "api_event",
		JobData: job,
	}
	client := GetClient("Test", "localhost:6379", "", 3)
	if err := client.SetEvent("coba-coba", 131*time.Second, dataEvent); err != nil {
		t.Fail()
	}
}
```
