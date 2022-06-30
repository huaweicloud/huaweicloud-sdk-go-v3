package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListDeploymentsRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	// 每页显示的条目数量，最大100，默认值10
	Limit *int64 `json:"limit,omitempty"`

	// 查询的起始位置，默认值0
	Offset *int64 `json:"offset,omitempty"`

	// 查询结果排序，如按照创建时间降序排序为created_at:desc，升序排序为created_at:asc
	Sort *string `json:"sort,omitempty"`

	// deployment名称（支持模糊匹配）
	Name *string `json:"name,omitempty"`

	// 节点ID，查询部署在该节点下的应用列表，和group_id不可同时请求
	NodeId *string `json:"node_id,omitempty"`

	// 节点组ID，查询部署在该节点组的应用列表，和node_id不可同时请求
	GroupId *string `json:"group_id,omitempty"`
}

func (o ListDeploymentsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDeploymentsRequest struct{}"
	}

	return strings.Join([]string{"ListDeploymentsRequest", string(data)}, " ")
}
