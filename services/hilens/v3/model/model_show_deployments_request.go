package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDeploymentsRequest Request Object
type ShowDeploymentsRequest struct {

	// 集群ID，查询部署在该节点组的应用列表，和node_id不可同时请求
	ClusterId *string `json:"cluster_id,omitempty"`

	// 节点ID，查询部署在该节点下的应用列表，和cluster_id不可同时请求
	NodeId *string `json:"node_id,omitempty"`

	// 平台提供者，分别为hilens及ief。当为hilens时，请求部署在hilens平台的相关数据
	Provider *string `json:"provider,omitempty"`

	// 部署名称(支持模糊匹配)
	Name *string `json:"name,omitempty"`

	// 查询结果排序，如按照创建时间降序排序为created_at:desc，升序排序为created_at:asc
	Sort *string `json:"sort,omitempty"`

	// 每页显示的条目数量, 最大 100，默认值 10
	Limit *int32 `json:"limit,omitempty"`

	// 查询的起始位置, 默认值 0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ShowDeploymentsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeploymentsRequest struct{}"
	}

	return strings.Join([]string{"ShowDeploymentsRequest", string(data)}, " ")
}
