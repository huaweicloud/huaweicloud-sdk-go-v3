package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowConfigurationDetailRequest struct {
	// 参数模板ID。

	ConfigId string `json:"config_id"`
}

func (o ShowConfigurationDetailRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowConfigurationDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowConfigurationDetailRequest", string(data)}, " ")
}
