package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateExternalEntityRequest Request Object
type CreateExternalEntityRequest struct {

	// 边缘节点ID
	EdgeNodeId string `json:"edge_node_id"`

	Body *CreateExternalEntityReqDto `json:"body,omitempty"`
}

func (o CreateExternalEntityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateExternalEntityRequest struct{}"
	}

	return strings.Join([]string{"CreateExternalEntityRequest", string(data)}, " ")
}
