package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchDeleteInstanceTopicRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`

	Body *BatchDeleteInstanceTopicReq `json:"body,omitempty"`
}

func (o BatchDeleteInstanceTopicRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchDeleteInstanceTopicRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteInstanceTopicRequest", string(data)}, " ")
}
