package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListApiQuantitiesV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
}

func (o ListApiQuantitiesV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListApiQuantitiesV2Request struct{}"
	}

	return strings.Join([]string{"ListApiQuantitiesV2Request", string(data)}, " ")
}
