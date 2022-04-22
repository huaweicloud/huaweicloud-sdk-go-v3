package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteSnapshotRequest struct {

	// 快照ID。
	SnapshotId string `json:"snapshot_id"`
}

func (o DeleteSnapshotRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSnapshotRequest struct{}"
	}

	return strings.Join([]string{"DeleteSnapshotRequest", string(data)}, " ")
}
