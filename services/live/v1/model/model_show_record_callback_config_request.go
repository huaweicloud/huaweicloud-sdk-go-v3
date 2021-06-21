package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowRecordCallbackConfigRequest struct {
	// 配置ID

	Id string `json:"id"`
}

func (o ShowRecordCallbackConfigRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowRecordCallbackConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowRecordCallbackConfigRequest", string(data)}, " ")
}
