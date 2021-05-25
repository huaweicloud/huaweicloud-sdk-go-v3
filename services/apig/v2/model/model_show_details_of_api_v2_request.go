package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowDetailsOfApiV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// API的编号

	ApiId string `json:"api_id"`
}

func (o ShowDetailsOfApiV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDetailsOfApiV2Request struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfApiV2Request", string(data)}, " ")
}
