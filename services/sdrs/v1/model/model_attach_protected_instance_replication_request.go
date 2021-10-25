package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type AttachProtectedInstanceReplicationRequest struct {
	// 保护实例的ID。

	ProtectedInstanceId string `json:"protected_instance_id"`

	Body *ProtectedInstanceAttachReplicationRequestBody `json:"body,omitempty"`
}

func (o AttachProtectedInstanceReplicationRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AttachProtectedInstanceReplicationRequest struct{}"
	}

	return strings.Join([]string{"AttachProtectedInstanceReplicationRequest", string(data)}, " ")
}
