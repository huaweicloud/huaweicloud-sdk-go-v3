package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowReplicationRequest struct {
	// 复制对ID。

	ReplicationId string `json:"replication_id"`
}

func (o ShowReplicationRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowReplicationRequest struct{}"
	}

	return strings.Join([]string{"ShowReplicationRequest", string(data)}, " ")
}
