package implementors

import (
	"net/http"
	"testing"
	"time"

	"git.bluebird.id/mybb/gorooster-client/v2/models"
)

var (
	client = GoroosterAPIImpl{
		BaseUrl:    "http://localhost:1407",
		ClientName: "ROOSTER-CLIENT-TEST",
		Client:     &http.Client{},
	}
	// dataRaw = map[string]interface{}{
	// 	"location": map[string]interface{}{
	// 		"latitude":  -6.246761122,
	// 		"longitude": 106.8256878,
	// 	},
	// 	"customer_id": "BB12345678",
	// 	"cid":         "0213",
	// }
	job = models.JobAPI{
		Endpoint: "https://jsonplaceholder.typicode.com/posts/1",
		Method:   models.METHOD_POST,
		Data:     nil,
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

	dataEvent = models.Event{
		Name:    "cancel_order_data",
		Id:      "101ec8dc-8de2-448c-b64c-6f0bc49cabff",
		Type:    models.API_EVENT,
		JobData: job,
	}
)

func TestGetEventApi(t *testing.T) {
	event := models.Event{
		JobData: models.JobAPI{},
	}
	client := GoroosterAPIImpl{
		BaseUrl:    "http://localhost:1407",
		ClientName: "ROOSTER-CLIENT-TEST",
		Client:     &http.Client{},
	}
	eventReleaseIn, err := client.GetEvent("CROW-10X", &event)
	if err != nil {
		t.Log(err)
	}
	t.Log(eventReleaseIn, event)
}
func TestCreateEventRelinApi(t *testing.T) {
	if err := client.SetEventReleaseIn("CROW-10X", 30*time.Hour, dataEvent); err != nil {
		t.Log(err.Error())
		t.Fail()
	}
}
func TestCreateEventRelatApi(t *testing.T) {
	if err := client.SetEventReleaseAt("CROW-11X", time.Now().Add(1*time.Hour), dataEvent); err != nil {
		t.Log(err.Error())
		t.Fail()
	}
}
func TestUpdateEventRelatApi(t *testing.T) {
	if err := client.UpdateReleaseEventAt("CROW-11X", time.Now().Add(1*time.Hour)); err != nil {
		t.Log(err.Error())
		t.Fail()
	}
}

func TestUpdateReleaseEventApi(t *testing.T) {
	client := GoroosterAPIImpl{
		BaseUrl:    "http://localhost:1407",
		ClientName: "ROOSTER-CLIENT-TEST",
		Client:     &http.Client{},
	}
	err := client.UpdateReleaseEventIn("CROW-10X", 10*time.Hour)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}
func TestUpdateDataEventApi(t *testing.T) {
	client := GoroosterAPIImpl{
		BaseUrl:    "http://localhost:1407",
		ClientName: "ROOSTER-CLIENT-TEST",
		Client:     &http.Client{},
	}
	dataRaw := map[string]interface{}{
		"location": map[string]interface{}{
			"latitude":  -6.246761122,
			"longitude": 106.8256878,
		},
		"customer_id": "BB12345678",
		"cid":         "0213",
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
		Name:    "cancel_order_data_X",
		Id:      "101ec8dc-8de2-448c-b64c-6f0bc49cabff",
		Type:    models.API_EVENT,
		JobData: job,
	}
	if err := client.UpdateDataEvent("CROW-10X", dataEvent); err != nil {
		t.Log(err.Error())
		t.Fail()
	}
}

func TestDeleteEventApi(t *testing.T) {
	client := GoroosterAPIImpl{
		BaseUrl:    "http://localhost:1407",
		ClientName: "ROOSTER-CLIENT-TEST",
		Client:     &http.Client{},
	}
	err := client.DeleteEvent("CROW-10X")
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}
