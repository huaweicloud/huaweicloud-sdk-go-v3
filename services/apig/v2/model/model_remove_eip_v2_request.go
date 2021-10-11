package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RemoveEipV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
}

func (o RemoveEipV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RemoveEipV2Request struct{}"
	}

	return strings.Join([]string{"RemoveEipV2Request", string(data)}, " ")
}
