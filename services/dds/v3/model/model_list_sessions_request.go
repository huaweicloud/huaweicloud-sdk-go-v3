package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSessionsRequest Request Object
type ListSessionsRequest struct {

	// 节点ID。允许查询的节点如下： 集群下面的 mongos节点以及 副本集、单节点实例下面的所有节点。
	NodeId string `json:"node_id"`

	// 索引位置，偏移量。 从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询记录数。取值范围[1, 20]，默认10 （表示返回10条数据）。
	Limit *int32 `json:"limit,omitempty"`

	// 执行计划描述。取空值表示查询所有语句类型，也可指定执行计划，例如： COLLSCAN IXSCAN FETCH SORT LIMIT SKIP COUNT COUNT_SCAN TEXT PROJECTION 等
	PlanSummary *string `json:"plan_summary,omitempty"`

	// 操作类型。取空值表示查询所有操作类型。也可指定操作类型，例如： none update insert query command getmore remove killcursors等
	Type *string `json:"type,omitempty"`

	// 命名空间。取空值表示查询所有命名空间。也可根据当前业务进行指定。
	Namespace *string `json:"namespace,omitempty"`

	// 运行时间，单位为 us。取空值表示查询所有的运行时间。也可根据当前业务需要进行配置，表示查询超出 cost_time 的会话。
	CostTime *int32 `json:"cost_time,omitempty"`
}

func (o ListSessionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSessionsRequest struct{}"
	}

	return strings.Join([]string{"ListSessionsRequest", string(data)}, " ")
}
