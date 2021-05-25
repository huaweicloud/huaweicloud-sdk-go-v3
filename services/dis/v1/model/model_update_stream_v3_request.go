package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateStreamV3Request struct {
	// 需要变更分区数量的通道名称。

	StreamName string `json:"stream_name"`

	Body *UpdateStreamRequest `json:"body,omitempty"`
}

func (o UpdateStreamV3Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateStreamV3Request struct{}"
	}

	return strings.Join([]string{"UpdateStreamV3Request", string(data)}, " ")
}
