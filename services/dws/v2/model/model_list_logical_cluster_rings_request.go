package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLogicalClusterRingsRequest Request Object
type ListLogicalClusterRingsRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 分页查询，偏移
	Offset *int32 `json:"offset,omitempty"`

	// 分页查询，每页显示的条目数量
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListLogicalClusterRingsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogicalClusterRingsRequest struct{}"
	}

	return strings.Join([]string{"ListLogicalClusterRingsRequest", string(data)}, " ")
}
