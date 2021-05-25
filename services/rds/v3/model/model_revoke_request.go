package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RevokeRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`
	// 实例ID。

	InstanceId string `json:"instance_id"`

	Body *RevokeRequestBody `json:"body,omitempty"`
}

func (o RevokeRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RevokeRequest struct{}"
	}

	return strings.Join([]string{"RevokeRequest", string(data)}, " ")
}
