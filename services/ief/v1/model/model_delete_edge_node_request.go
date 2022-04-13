package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteEdgeNodeRequest struct {
	// 边缘节点ID

	NodeId string `json:"node_id"`
	// 铂金版实例ID，专业版实例为空值

	IefInstanceId *string `json:"ief-instance-id,omitempty"`
}

func (o DeleteEdgeNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEdgeNodeRequest struct{}"
	}

	return strings.Join([]string{"DeleteEdgeNodeRequest", string(data)}, " ")
}
