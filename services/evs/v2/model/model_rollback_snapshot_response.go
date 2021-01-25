package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RollbackSnapshotResponse struct {
	Rollback       *RollbackInfo `json:"rollback,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o RollbackSnapshotResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RollbackSnapshotResponse struct{}"
	}

	return strings.Join([]string{"RollbackSnapshotResponse", string(data)}, " ")
}
