package helpers

import (
	"fmt"
	"strings"
)

func GenerateKeyEvent(clientName, key string) string {
	keyEvent := fmt.Sprintf("%s:event:%s", clientName, key)
	return keyEvent
}

func GenerateKeyData(clientName, key string) string {
	keyData := fmt.Sprintf("%s:data:%s", clientName, key)
	return keyData
}

func ValidatorClientNameAndKey(str string) bool {
	return !strings.Contains(str, ":")

}

type SuccessResponse struct {
	Event          interface{} `json:"event"`
	EventReleaseIn string      `json:"event_release_in"`
}

// ================================ error format ================================
type Error struct {
	StatusCode       int                    `json:"-"`
	ErrorCode        string                 `json:"error_code"`
	ErrorMessage     string                 `json:"error_message"`
	ErrorField       string                 `json:"error_field,omitempty"`
	LocalizedMessage Message                `json:"localized_message"`
	Data             map[string]interface{} `json:"data,omitempty"`
	ErrorData        interface{}            `json:"error_data,omitempty"`
}
type Message struct {
	English   string `json:"en"`
	Indonesia string `json:"id"`
}
