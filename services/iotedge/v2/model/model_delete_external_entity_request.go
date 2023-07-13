package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteExternalEntityRequest Request Object
type DeleteExternalEntityRequest struct {

	// 边缘节点ID
	EdgeNodeId string `json:"edge_node_id"`

	// 外部实体ID
	ExternalId string `json:"external_id"`
}

func (o DeleteExternalEntityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteExternalEntityRequest struct{}"
	}

	return strings.Join([]string{"DeleteExternalEntityRequest", string(data)}, " ")
}
