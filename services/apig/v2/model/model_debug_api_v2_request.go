package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DebugApiV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// API的编号

	ApiId string `json:"api_id"`

	Body *DebugApiReq `json:"body,omitempty"`
}

func (o DebugApiV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DebugApiV2Request struct{}"
	}

	return strings.Join([]string{"DebugApiV2Request", string(data)}, " ")
}
