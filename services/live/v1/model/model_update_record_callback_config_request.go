package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateRecordCallbackConfigRequest struct {
	// 配置ID，在创建配置成功后返回

	Id string `json:"id"`

	Body *RecordCallbackConfigRequest `json:"body,omitempty"`
}

func (o UpdateRecordCallbackConfigRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateRecordCallbackConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateRecordCallbackConfigRequest", string(data)}, " ")
}
