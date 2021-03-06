package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteApiGroupV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// 分组的编号

	GroupId string `json:"group_id"`
}

func (o DeleteApiGroupV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteApiGroupV2Request struct{}"
	}

	return strings.Join([]string{"DeleteApiGroupV2Request", string(data)}, " ")
}
