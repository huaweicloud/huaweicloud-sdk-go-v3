package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSnapshotRequest Request Object
type ShowSnapshotRequest struct {

	// 快照ID。
	SnapshotId string `json:"snapshot_id"`
}

func (o ShowSnapshotRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSnapshotRequest struct{}"
	}

	return strings.Join([]string{"ShowSnapshotRequest", string(data)}, " ")
}
