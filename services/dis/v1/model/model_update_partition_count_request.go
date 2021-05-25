package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdatePartitionCountRequest struct {
	// 需要变更分区数量的通道名称。

	StreamName string `json:"stream_name"`

	Body *UpdatePartitionCountRequestBody `json:"body,omitempty"`
}

func (o UpdatePartitionCountRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdatePartitionCountRequest struct{}"
	}

	return strings.Join([]string{"UpdatePartitionCountRequest", string(data)}, " ")
}
