package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListEdgeNodeCertsRequest struct {
	// 节点ID

	NodeId string `json:"node_id"`
	// 铂金版实例ID，专业版实例为空值

	IefInstanceId *string `json:"ief-instance-id,omitempty"`
	// 查询返回记录的数量限制

	Limit *int32 `json:"limit,omitempty"`
	// 偏移量，表示查询该偏移量后面的记录

	Offset *int32 `json:"offset,omitempty"`
}

func (o ListEdgeNodeCertsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEdgeNodeCertsRequest struct{}"
	}

	return strings.Join([]string{"ListEdgeNodeCertsRequest", string(data)}, " ")
}
