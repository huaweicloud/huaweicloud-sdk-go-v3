package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ChangeApiVersionV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// API的编号

	ApiId string `json:"api_id"`

	Body *ApiVersionInfo `json:"body,omitempty"`
}

func (o ChangeApiVersionV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ChangeApiVersionV2Request struct{}"
	}

	return strings.Join([]string{"ChangeApiVersionV2Request", string(data)}, " ")
}
