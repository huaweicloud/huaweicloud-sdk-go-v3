package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateInstanceConfigurationRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`
	// 实例ID。

	InstanceId string `json:"instance_id"`

	Body *OpenGaussModifyInstanceConfigurationRequest `json:"body,omitempty"`
}

func (o UpdateInstanceConfigurationRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateInstanceConfigurationRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceConfigurationRequest", string(data)}, " ")
}
