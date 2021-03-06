package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CopyCheckpointResponse struct {
	Replication    *CheckpointReplicateRespBody `json:"replication,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o CopyCheckpointResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CopyCheckpointResponse struct{}"
	}

	return strings.Join([]string{"CopyCheckpointResponse", string(data)}, " ")
}
