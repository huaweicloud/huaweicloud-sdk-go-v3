package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDsRequest Request Object
type CreateDsRequest struct {

	// 边缘节点ID
	EdgeNodeId string `json:"edge_node_id"`

	Body *CreateDcDsReqDto `json:"body,omitempty"`
}

func (o CreateDsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDsRequest struct{}"
	}

	return strings.Join([]string{"CreateDsRequest", string(data)}, " ")
}
