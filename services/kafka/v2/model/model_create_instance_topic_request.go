package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateInstanceTopicRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`

	Body *CreateInstanceTopicReq `json:"body,omitempty"`
}

func (o CreateInstanceTopicRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateInstanceTopicRequest struct{}"
	}

	return strings.Join([]string{"CreateInstanceTopicRequest", string(data)}, " ")
}
