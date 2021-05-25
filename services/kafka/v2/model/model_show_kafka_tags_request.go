package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowKafkaTagsRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`
}

func (o ShowKafkaTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowKafkaTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowKafkaTagsRequest", string(data)}, " ")
}
