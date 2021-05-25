package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowInstanceStatusRequest struct {
	// 语言类型 中文:zh-cn 英文:en-us，默认en-us

	XLanguage *string `json:"X-Language,omitempty"`
	// 实例ID

	TaskId string `json:"task_id"`
}

func (o ShowInstanceStatusRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowInstanceStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceStatusRequest", string(data)}, " ")
}
