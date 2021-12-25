package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateEdgeNodeCertsRequest struct {
	// 边缘节点ID

	NodeId string `json:"node_id"`
	// 铂金版实例ID，专业版实例为空值

	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	Body *Cert `json:"body,omitempty"`
}

func (o CreateEdgeNodeCertsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEdgeNodeCertsRequest struct{}"
	}

	return strings.Join([]string{"CreateEdgeNodeCertsRequest", string(data)}, " ")
}
