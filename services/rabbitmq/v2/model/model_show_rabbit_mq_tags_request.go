package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowRabbitMqTagsRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`
}

func (o ShowRabbitMqTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowRabbitMqTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowRabbitMqTagsRequest", string(data)}, " ")
}
