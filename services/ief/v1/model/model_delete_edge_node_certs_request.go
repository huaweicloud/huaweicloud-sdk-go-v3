package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteEdgeNodeCertsRequest struct {
	// 边缘节点ID

	NodeId string `json:"node_id"`
	// 证书id

	CertId string `json:"cert_id"`
	// 铂金版实例ID，专业版实例为空值

	IefInstanceId *string `json:"ief-instance-id,omitempty"`
}

func (o DeleteEdgeNodeCertsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEdgeNodeCertsRequest struct{}"
	}

	return strings.Join([]string{"DeleteEdgeNodeCertsRequest", string(data)}, " ")
}
