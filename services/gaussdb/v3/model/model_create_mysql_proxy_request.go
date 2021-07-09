package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateMysqlProxyRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`
	// 实例ID

	InstanceId string `json:"instance_id"`

	Body *OpenMysqlProxyRequestBody `json:"body,omitempty"`
}

func (o CreateMysqlProxyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateMysqlProxyRequest struct{}"
	}

	return strings.Join([]string{"CreateMysqlProxyRequest", string(data)}, " ")
}
