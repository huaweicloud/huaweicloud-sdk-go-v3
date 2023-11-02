package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSnapshotCrossRegionPolicyRequest Request Object
type ListSnapshotCrossRegionPolicyRequest struct {

	// 集群ID
	ClusterId *string `json:"cluster_id,omitempty"`

	// 分页偏移
	Offset *int32 `json:"offset,omitempty"`

	// 分页大小
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListSnapshotCrossRegionPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSnapshotCrossRegionPolicyRequest struct{}"
	}

	return strings.Join([]string{"ListSnapshotCrossRegionPolicyRequest", string(data)}, " ")
}
