package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchCreateOrDeleteKafkaTagRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`

	Body *BatchCreateOrDeleteTagReq `json:"body,omitempty"`
}

func (o BatchCreateOrDeleteKafkaTagRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchCreateOrDeleteKafkaTagRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateOrDeleteKafkaTagRequest", string(data)}, " ")
}
