package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateInstanceNameRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`
	// 实例ID，严格匹配UUID规则。

	InstanceId string `json:"instance_id"`

	Body *UpdateNameRequestBody `json:"body,omitempty"`
}

func (o UpdateInstanceNameRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateInstanceNameRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceNameRequest", string(data)}, " ")
}
