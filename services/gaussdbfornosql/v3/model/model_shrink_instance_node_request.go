package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShrinkInstanceNodeRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`

	Body *ShrinkInstanceNodeRequestBody `json:"body,omitempty"`
}

func (o ShrinkInstanceNodeRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShrinkInstanceNodeRequest struct{}"
	}

	return strings.Join([]string{"ShrinkInstanceNodeRequest", string(data)}, " ")
}
