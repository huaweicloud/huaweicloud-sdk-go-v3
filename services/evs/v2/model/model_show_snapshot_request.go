package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowSnapshotRequest struct {
	// 快照ID。

	SnapshotId string `json:"snapshot_id"`
}

func (o ShowSnapshotRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowSnapshotRequest struct{}"
	}

	return strings.Join([]string{"ShowSnapshotRequest", string(data)}, " ")
}
