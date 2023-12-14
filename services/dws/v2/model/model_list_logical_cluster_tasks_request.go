package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLogicalClusterTasksRequest Request Object
type ListLogicalClusterTasksRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 分页查询，偏移
	Offset *int32 `json:"offset,omitempty"`

	// 分页查询，每页显示的条目数量
	Limit *int32 `json:"limit,omitempty"`

	// 集群名称
	LogicalClusterName *string `json:"logical_cluster_name,omitempty"`

	// 类型
	Type *string `json:"type,omitempty"`

	// 排序字段
	OrderBy *string `json:"order_by,omitempty"`

	// 排序：升序/降序
	Order *string `json:"order,omitempty"`
}

func (o ListLogicalClusterTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogicalClusterTasksRequest struct{}"
	}

	return strings.Join([]string{"ListLogicalClusterTasksRequest", string(data)}, " ")
}
