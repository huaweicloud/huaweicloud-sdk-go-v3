package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLogicalClustersRequest Request Object
type ListLogicalClustersRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 分页查询，偏移
	Offset *int32 `json:"offset,omitempty"`

	// 分页查询，每页显示的条目数量
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListLogicalClustersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogicalClustersRequest struct{}"
	}

	return strings.Join([]string{"ListLogicalClustersRequest", string(data)}, " ")
}
