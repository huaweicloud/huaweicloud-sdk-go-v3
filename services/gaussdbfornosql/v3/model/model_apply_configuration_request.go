package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ApplyConfigurationRequest struct {
	// 参数模板ID。

	ConfigId string `json:"config_id"`

	Body *ApplyConfigurationRequestBody `json:"body,omitempty"`
}

func (o ApplyConfigurationRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ApplyConfigurationRequest struct{}"
	}

	return strings.Join([]string{"ApplyConfigurationRequest", string(data)}, " ")
}
