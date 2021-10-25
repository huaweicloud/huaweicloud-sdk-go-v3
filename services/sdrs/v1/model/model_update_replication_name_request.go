package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateReplicationNameRequest struct {
	// replication_id

	ReplicationId string `json:"replication_id"`

	Body *UpdateReplicationNameRequestBody `json:"body,omitempty"`
}

func (o UpdateReplicationNameRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateReplicationNameRequest struct{}"
	}

	return strings.Join([]string{"UpdateReplicationNameRequest", string(data)}, " ")
}
