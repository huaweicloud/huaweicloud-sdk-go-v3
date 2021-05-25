package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowConfigurationRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`
	// 参数模板ID。

	ConfigId string `json:"config_id"`
}

func (o ShowConfigurationRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowConfigurationRequest struct{}"
	}

	return strings.Join([]string{"ShowConfigurationRequest", string(data)}, " ")
}
