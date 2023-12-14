package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterNodesRequest Request Object
type ListClusterNodesRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 是否被删除，true/false
	Deleted *string `json:"deleted,omitempty"`

	// 节点ID列表
	NodeIds *[]string `json:"node_ids,omitempty"`

	// 分页查询，偏移
	Offset *int32 `json:"offset,omitempty"`

	// 分页查询，每页显示的条目数量
	Limit *int32 `json:"limit,omitempty"`

	// 过滤字段
	FilterBy *string `json:"filter_by,omitempty"`

	// 过滤字段内容
	Filter *string `json:"filter,omitempty"`

	// 排序字段
	OrderBy *string `json:"order_by,omitempty"`

	// 排序：升序/降序
	Order *string `json:"order,omitempty"`
}

func (o ListClusterNodesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterNodesRequest struct{}"
	}

	return strings.Join([]string{"ListClusterNodesRequest", string(data)}, " ")
}
