package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateSnapshotRequest struct {

	// 快照ID
	SnapshotId string `json:"snapshot_id"`

	Body *UpdateSnapshotRequestBody `json:"body,omitempty"`
}

func (o UpdateSnapshotRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSnapshotRequest struct{}"
	}

	return strings.Join([]string{"UpdateSnapshotRequest", string(data)}, " ")
}
