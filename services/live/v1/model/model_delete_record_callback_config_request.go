package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteRecordCallbackConfigRequest struct {
	// 配置ID

	Id string `json:"id"`
}

func (o DeleteRecordCallbackConfigRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteRecordCallbackConfigRequest struct{}"
	}

	return strings.Join([]string{"DeleteRecordCallbackConfigRequest", string(data)}, " ")
}
