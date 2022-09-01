package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateExternalEntityRequest struct {

	// 边缘节点ID
	EdgeNodeId string `json:"edge_node_id" xml:"edge_node_id"`

	// 外部实体ID
	ExternalId string `json:"external_id" xml:"external_id"`

	Body *UpdateExternalEntityReqDto `json:"body,omitempty" xml:"body"`
}

func (o UpdateExternalEntityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateExternalEntityRequest struct{}"
	}

	return strings.Join([]string{"UpdateExternalEntityRequest", string(data)}, " ")
}
