package client

import (
	"testing"
	"time"

	"github.com/LukmanulHakim18/gorooster-client/models"
)

var (
	key        = "202ec8dc-8de2-448c-b64c-6f0bc49cabff"
	redisHost  = "localhost:6379"
	Pass       = ""
	DB         = 14
	ClientName = "rooster_client_test"
)

func TestSetEvent(t *testing.T) {
	dataRaw := map[string]interface{}{
		"location": map[string]interface{}{
			"latitude":  -6.246761122,
			"longitude": 106.8256878,
		},
		"customer_id": "BB12345678",
		"cid":         "021",
	}
	job := models.JobAPI{
		Endpoint: "https://foo.id/bar",
		Method:   models.METHOD_POST,
		Data:     dataRaw,
		Headers: []models.Header{
			{
				Key:   "Token",
				Value: "my-token",
			},
			{
				Key:   "Content-Type",
				Value: "application/json",
			},
		},
	}

	dataEvent := models.Event{
		Name:    "cancel_order",
		Id:      "101ec8dc-8de2-448c-b64c-6f0bc49cabff",
		Type:    models.API_EVENT,
		JobData: job,
	}
	client := GetRedisClient(ClientName, redisHost, Pass, DB)
	if err := client.SetEvent(key, 100*time.Minute, dataEvent); err != nil {
		t.Fail()
	}
}

func TestGetEvent(t *testing.T) {
	client := GetRedisClient(ClientName, redisHost, Pass, DB)
	event := models.Event{
		JobData: models.JobAPI{},
	}
	ttl, err := client.GetEvent(key, &event)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	t.Log(ttl, event)
}
func TestUpdateEventRelease(t *testing.T) {
	client := GetRedisClient(ClientName, redisHost, Pass, DB)
	if err := client.UpdateReleaseEvent(key, 10*time.Second); err != nil {
		t.Log(err.Error())
		t.Fail()
	}
}

func TestUpdateDataEvent(t *testing.T) {
	dataRaw := map[string]interface{}{
		"location": map[string]interface{}{
			"latitude":  -6.246761122,
			"longitude": 16.8256878,
		},
		"customer_id": "BB12345678",
		"cid":         "021",
	}
	job := models.JobAPI{
		Endpoint: "https://stgapi4.bluebird.id/odrd/trip/newid",
		Method:   models.METHOD_POST,
		Data:     dataRaw,
		Headers: []models.Header{
			{
				Key:   "Token",
				Value: "mybb-odrd-token",
			},
			{
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
	client := GetRedisClient(ClientName, redisHost, Pass, DB)
	if err := client.UpdateDataEvent(key, dataEvent); err != nil {
		t.Fail()
	}
}

func TestDeleteEvent(t *testing.T) {
	client := GetRedisClient(ClientName, redisHost, Pass, DB)
	if err := client.DeleteEvent(key); err != nil {
		t.Log(err.Error())
		t.Fail()
	}
}

func TestErrorClinetName(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	GetRedisClient("must:error", redisHost, Pass, DB)

}
func TestErrorKey(t *testing.T) {

	client := GetRedisClient(ClientName, redisHost, Pass, DB)
	if err := client.SetEvent("event:error", 10*time.Second, models.Event{}); err == nil {
		t.Fail()
	}

}
