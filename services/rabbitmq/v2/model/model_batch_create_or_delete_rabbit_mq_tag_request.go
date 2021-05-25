package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchCreateOrDeleteRabbitMqTagRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`

	Body *BatchCreateOrDeleteTagReq `json:"body,omitempty"`
}

func (o BatchCreateOrDeleteRabbitMqTagRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchCreateOrDeleteRabbitMqTagRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateOrDeleteRabbitMqTagRequest", string(data)}, " ")
}
