package implementors

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"git.bluebird.id/mybb/gorooster-client/v2/helpers"
	"git.bluebird.id/mybb/gorooster-client/v2/models"
)

type GoroosterAPIImpl struct {
	Client     *http.Client
	APIKey     string
	ClientName string
	BaseUrl    string
}

func (impl GoroosterAPIImpl) GetEvent(key string, target interface{}) (eventReleaseIn time.Duration, err error) {
	endponit := fmt.Sprintf("%s/event/%s", impl.BaseUrl, key)
	body := bytes.NewReader(nil)
	req, err := http.NewRequest("GET", endponit, body)
	if err != nil {
		return
	}
	res, err := impl.runRequest(req)
	if err != nil {
		return
	}
	responseData := helpers.SuccessResponse{Event: target}
	if err = json.NewDecoder(res.Body).Decode(&responseData); err != nil {
		return
	}
	if eventReleaseIn, err = time.ParseDuration(responseData.EventReleaseIn); err != nil {
		return
	}
	return
}

func (impl GoroosterAPIImpl) SetEvent(key string, eventReleaseIn time.Duration, event models.Event) error {
	endponit := fmt.Sprintf("%s/event/%s/%s", impl.BaseUrl, key, eventReleaseIn.String())
	payload, err := json.Marshal(event)
	if err != nil {
		return err
	}
	body := bytes.NewReader(payload)
	req, err := http.NewRequest("POST", endponit, body)
	if err != nil {
		return err
	}
	if _, err := impl.runRequest(req); err != nil {
		return err
	}

	return nil
}

func (impl GoroosterAPIImpl) UpdateReleaseEvent(key string, eventReleaseIn time.Duration) (err error) {
	endponit := fmt.Sprintf("%s/event/%s/%s", impl.BaseUrl, key, eventReleaseIn.String())
	body := bytes.NewReader(nil)
	req, err := http.NewRequest("PUT", endponit, body)
	if err != nil {
		return
	}
	_, err = impl.runRequest(req)
	if err != nil {
		return
	}
	return nil
}

func (impl GoroosterAPIImpl) UpdateDataEvent(key string, event models.Event) (err error) {

	endponit := fmt.Sprintf("%s/event/%s", impl.BaseUrl, key)
	payload, err := json.Marshal(event)
	if err != nil {
		return err
	}
	body := bytes.NewReader(payload)
	req, err := http.NewRequest("PUT", endponit, body)
	if err != nil {
		return
	}
	_, err = impl.runRequest(req)
	if err != nil {
		return
	}
	return nil

}

func (impl GoroosterAPIImpl) DeleteEvent(key string) (err error) {

	endponit := fmt.Sprintf("%s/event/%s", impl.BaseUrl, key)
	body := bytes.NewReader(nil)
	req, err := http.NewRequest("DELETE", endponit, body)
	if err != nil {
		return
	}
	_, err = impl.runRequest(req)
	if err != nil {
		return
	}
	return
}

func (impl GoroosterAPIImpl) runRequest(req *http.Request) (res *http.Response, err error) {

	req.Header.Set("X-CLIENT-NAME", impl.ClientName)
	req.Header.Set("Content-Type", "application/json")
	res, err = impl.Client.Do(req)
	if err != nil {
		return
	}
	err = CekResponseHttp(res)
	if err != nil {
		return
	}
	return
}

func CekResponseHttp(res *http.Response) (err error) {
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		StdError := helpers.Error{}
		err = json.NewDecoder(res.Body).Decode(&StdError)
		if err != nil {
			return
		}
		return fmt.Errorf("status_code: %d, endpoint: %s, message: %s", res.StatusCode, res.Request.URL, StdError.ErrorMessage)
	}
	return nil

}

func (impl GoroosterAPIImpl) SetEventAt(key string, eventReleaseAt time.Time, event models.Event) error {
	endponit := fmt.Sprintf("%s/event/release_at/%s/%d", impl.BaseUrl, key, eventReleaseAt.Unix())
	payload, err := json.Marshal(event)
	if err != nil {
		return err
	}
	body := bytes.NewReader(payload)
	req, err := http.NewRequest("POST", endponit, body)
	if err != nil {
		return err
	}
	if _, err := impl.runRequest(req); err != nil {
		return err
	}

	return nil
}

func (impl GoroosterAPIImpl) UpdateReleaseEventAt(key string, eventReleaseAt time.Time) (err error) {
	endponit := fmt.Sprintf("%s/event/release_at/%s/%d", impl.BaseUrl, key, eventReleaseAt.Unix())
	body := bytes.NewReader(nil)
	req, err := http.NewRequest("PUT", endponit, body)
	if err != nil {
		return
	}
	_, err = impl.runRequest(req)
	if err != nil {
		return
	}
	return nil
}
