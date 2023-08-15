package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterSnapshotsRequest Request Object
type ListClusterSnapshotsRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 查询条数
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 排序字段
	SortKey *string `json:"sort_key,omitempty"`

	// 排序规则
	SortDir *string `json:"sort_dir,omitempty"`
}

func (o ListClusterSnapshotsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterSnapshotsRequest struct{}"
	}

	return strings.Join([]string{"ListClusterSnapshotsRequest", string(data)}, " ")
}
