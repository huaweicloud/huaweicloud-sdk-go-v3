package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateExternalEntityRequest struct {
	// 边缘节点ID

	EdgeNodeId string `json:"edge_node_id"`
	// 外部实体ID

	ExternalId string `json:"external_id"`

	Body *UpdateExternalEntityReqDto `json:"body,omitempty"`
}

func (o UpdateExternalEntityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateExternalEntityRequest struct{}"
	}

	return strings.Join([]string{"UpdateExternalEntityRequest", string(data)}, " ")
}
