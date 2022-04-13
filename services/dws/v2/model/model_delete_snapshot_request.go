package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteSnapshotRequest struct {
	// 待删除的快照ID

	SnapshotId string `json:"snapshot_id"`
}

func (o DeleteSnapshotRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSnapshotRequest struct{}"
	}

	return strings.Join([]string{"DeleteSnapshotRequest", string(data)}, " ")
}
