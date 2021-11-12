package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListExternalEntityRequest struct {
	// 边缘节点ID

	EdgeNodeId string `json:"edge_node_id"`
	// 查询的起始位置，取值范围为非负整数，默认为0

	Offset *int32 `json:"offset,omitempty"`
	// 每页记录数，取值范围为非负整数，默认值为10

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListExternalEntityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListExternalEntityRequest struct{}"
	}

	return strings.Join([]string{"ListExternalEntityRequest", string(data)}, " ")
}
