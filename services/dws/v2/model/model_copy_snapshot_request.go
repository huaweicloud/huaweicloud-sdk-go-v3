package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CopySnapshotRequest struct {

	// 快照ID
	SnapshotId string `json:"snapshot_id"`

	Body *interface{} `json:"body,omitempty"`
}

func (o CopySnapshotRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopySnapshotRequest struct{}"
	}

	return strings.Join([]string{"CopySnapshotRequest", string(data)}, " ")
}
