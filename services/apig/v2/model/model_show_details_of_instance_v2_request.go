package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowDetailsOfInstanceV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
}

func (o ShowDetailsOfInstanceV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDetailsOfInstanceV2Request struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfInstanceV2Request", string(data)}, " ")
}
