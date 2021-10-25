package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateReplicationRequest struct {
	Body *CreateReplicationRequestBody `json:"body,omitempty"`
}

func (o CreateReplicationRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateReplicationRequest struct{}"
	}

	return strings.Join([]string{"CreateReplicationRequest", string(data)}, " ")
}
