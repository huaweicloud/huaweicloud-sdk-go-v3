package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ExpandReplicationRequest struct {
	// 复制对的ID。

	ReplicationId string `json:"replication_id"`

	Body *ExtendReplicationRequestBody `json:"body,omitempty"`
}

func (o ExpandReplicationRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ExpandReplicationRequest struct{}"
	}

	return strings.Join([]string{"ExpandReplicationRequest", string(data)}, " ")
}
