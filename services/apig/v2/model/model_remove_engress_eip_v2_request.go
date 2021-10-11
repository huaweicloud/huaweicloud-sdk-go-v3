package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RemoveEngressEipV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
}

func (o RemoveEngressEipV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RemoveEngressEipV2Request struct{}"
	}

	return strings.Join([]string{"RemoveEngressEipV2Request", string(data)}, " ")
}
