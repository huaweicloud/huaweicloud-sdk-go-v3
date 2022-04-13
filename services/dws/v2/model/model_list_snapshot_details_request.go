package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListSnapshotDetailsRequest struct {
	// 快照ID

	SnapshotId string `json:"snapshot_id"`
}

func (o ListSnapshotDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSnapshotDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListSnapshotDetailsRequest", string(data)}, " ")
}
