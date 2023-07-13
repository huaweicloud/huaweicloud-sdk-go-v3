package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RollbackSnapshotRequest Request Object
type RollbackSnapshotRequest struct {

	// 快照ID
	SnapshotId string `json:"snapshot_id"`

	Body *RollbackSnapshotRequestBody `json:"body,omitempty"`
}

func (o RollbackSnapshotRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RollbackSnapshotRequest struct{}"
	}

	return strings.Join([]string{"RollbackSnapshotRequest", string(data)}, " ")
}
