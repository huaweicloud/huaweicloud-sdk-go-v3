package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEdgeNodeRequest Request Object
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
