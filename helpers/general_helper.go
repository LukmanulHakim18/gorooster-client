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

func ValidatorClinetNameAndKey(str string) bool {
	return !strings.Contains(str, ":")

}
