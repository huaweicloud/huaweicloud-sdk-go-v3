package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteMysqlProxyRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`
	// 实例ID

	InstanceId string `json:"instance_id"`
}

func (o DeleteMysqlProxyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteMysqlProxyRequest struct{}"
	}

	return strings.Join([]string{"DeleteMysqlProxyRequest", string(data)}, " ")
}
