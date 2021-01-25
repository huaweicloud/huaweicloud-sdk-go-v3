package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CopyCheckpointRequest struct {
	Body *CheckpointReplicateReq `json:"body,omitempty"`
}

func (o CopyCheckpointRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CopyCheckpointRequest struct{}"
	}

	return strings.Join([]string{"CopyCheckpointRequest", string(data)}, " ")
}
