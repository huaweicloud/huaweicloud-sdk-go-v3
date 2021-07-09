package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowMysqlProxyRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`
	// 实例ID

	InstanceId string `json:"instance_id"`
}

func (o ShowMysqlProxyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowMysqlProxyRequest struct{}"
	}

	return strings.Join([]string{"ShowMysqlProxyRequest", string(data)}, " ")
}
