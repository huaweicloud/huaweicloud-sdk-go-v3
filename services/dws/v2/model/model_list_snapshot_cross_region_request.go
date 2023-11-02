package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSnapshotCrossRegionRequest Request Object
type ListSnapshotCrossRegionRequest struct {

	// 分页偏移，默认0
	Offset *int32 `json:"offset,omitempty"`

	// 分页大小，默认10
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListSnapshotCrossRegionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSnapshotCrossRegionRequest struct{}"
	}

	return strings.Join([]string{"ListSnapshotCrossRegionRequest", string(data)}, " ")
}
