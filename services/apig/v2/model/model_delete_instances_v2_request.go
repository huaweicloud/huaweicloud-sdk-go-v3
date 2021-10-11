package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteInstancesV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
}

func (o DeleteInstancesV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteInstancesV2Request struct{}"
	}

	return strings.Join([]string{"DeleteInstancesV2Request", string(data)}, " ")
}
