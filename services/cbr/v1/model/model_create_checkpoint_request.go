package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateCheckpointRequest struct {
	Body *VaultBackupReq `json:"body,omitempty"`
}

func (o CreateCheckpointRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateCheckpointRequest struct{}"
	}

	return strings.Join([]string{"CreateCheckpointRequest", string(data)}, " ")
}
