package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListApiGroupsQuantitiesV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
}

func (o ListApiGroupsQuantitiesV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListApiGroupsQuantitiesV2Request struct{}"
	}

	return strings.Join([]string{"ListApiGroupsQuantitiesV2Request", string(data)}, " ")
}
