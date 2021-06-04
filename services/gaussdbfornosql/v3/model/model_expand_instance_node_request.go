package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ExpandInstanceNodeRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`

	Body *ExpandInstanceNodeRequestBody `json:"body,omitempty"`
}

func (o ExpandInstanceNodeRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ExpandInstanceNodeRequest struct{}"
	}

	return strings.Join([]string{"ExpandInstanceNodeRequest", string(data)}, " ")
}
