package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowMysqlProxyFlavorsRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`
	// 实例ID

	InstanceId string `json:"instance_id"`
}

func (o ShowMysqlProxyFlavorsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowMysqlProxyFlavorsRequest struct{}"
	}

	return strings.Join([]string{"ShowMysqlProxyFlavorsRequest", string(data)}, " ")
}
