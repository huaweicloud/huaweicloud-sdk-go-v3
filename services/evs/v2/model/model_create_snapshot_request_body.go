package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type CreateSnapshotRequestBody struct {
	Snapshot *CreateSnapshotOption `json:"snapshot"`
}

func (o CreateSnapshotRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateSnapshotRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSnapshotRequestBody", string(data)}, " ")
}
