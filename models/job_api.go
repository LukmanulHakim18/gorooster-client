package models

import (
	"fmt"
	"strings"
)

type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type MethodType string

const (
	METHOD_POST   = "POST"
	METHOD_GET    = "GET"
	METHOD_PUT    = "PUT"
	METHOD_PATCH  = "PATCH"
	METHOD_DELETE = "DELETE"
)

type JobAPI struct {
	Endpoint string      `json:"endpoint"`
	Headers  []Header    `json:"headers"`
	Method   MethodType  `json:"method"`
	Data     interface{} `json:"data"`
}

// Check if methode support or not
func (ja JobAPI) IsMethodSupport() bool {
	return ja.Method == METHOD_POST || ja.Method == METHOD_GET || ja.Method == METHOD_PUT || ja.Method == METHOD_DELETE || ja.Method == METHOD_PATCH
}

// parsing  methode to string
func (m MethodType) ToString() string {
	return string(m)
}

// Validate data Job API
func (ja *JobAPI) Validate() error {
	methodeType := strings.ToUpper(ja.Method.ToString())
	ja.Method = MethodType(methodeType)
	if ok := ja.IsMethodSupport(); !ok {
		return fmt.Errorf("method not support")
	}
	if ja.Endpoint == "" {
		return fmt.Errorf("endpoint is empty")

	}
	return nil
}
