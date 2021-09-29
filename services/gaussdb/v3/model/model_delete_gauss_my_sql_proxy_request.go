package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteGaussMySqlProxyRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`
	// 实例ID

	InstanceId string `json:"instance_id"`
}

func (o DeleteGaussMySqlProxyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteGaussMySqlProxyRequest struct{}"
	}

	return strings.Join([]string{"DeleteGaussMySqlProxyRequest", string(data)}, " ")
}
