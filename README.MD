# Gorooster-client

This repo use for support event scheduler service [Gorooster](https://github.com/LukmanulHakim18/gorooster) for set event scheduler.

# Features

- Set event release In
- Set event release At
- Update release event in
- Update release event an
- Get event
- Update data event
- Delete event

# Installation

Gorooster requires a Go version with modules support. So make sure to initialize a Go module:

```go
go mod init github.com/my/repo
```

And instal the lib

```go
go get github.com/LukmanulHakim18/gorooster-client
```

# Quick start

## connection via rest-api

```go
import (
	"fmt"
	"time"

	client "github.com/LukmanulHakim18/gorooster-client"
	clientModels "github.com/LukmanulHakim18/gorooster-client/models"
)

func ExampleClient() {
	clientName := "rooster_weakup" // client name can't contains ':'
	gc := client.GetAPIClient("Timesheet", "http://localhost:1407","")
	dataRaw := map[string]interface{}{
		"location": map[string]interface{}{
			"latitude":  -6.246761122,
			"longitude": 106.8256878,
		},
		"customer_id": "BB12345678",
		"cid":         "021",
	}
	job := clientModels.JobAPI{
		Endpoint: "https://gorooster.id/foo/bar",
		Method:   clientModels.METHOD_POST,
		Data:     dataRaw,
		Headers: []clientModels.Header{
			clientModels.Header{
				Key:   "Token",
				Value: "my-token",
			},
			clientModels.Header{
				Key:   "Content-Type",
				Value: "application/json",
			},
		},
	}

	dataEvent := clientModels.Event{
		Name:    "cancel_order",
		Id:      "101ec8dc-8de2-448c-b64c-6f0bc49cabff",
		Type:    "api_event",
		JobData: job,
	}
	eventName:="test-event-1"  // event name can't contains ':' and must uniq
	err := gc.SetEventReleaseIn(eventName, 10*time.Second, dataEvent)
	if err != nil {
		fmt.Print(err.Error())
	}
}
```

# Run the test

Make sure your redis started and your config was correct

Just run comment below:

```go
go test
```
