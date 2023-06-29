package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEdgeNodesResponse Response Object
type ListEdgeNodesResponse struct {

	// 边缘节点列表
	Nodes *[]EdgeNodeResp `json:"nodes,omitempty"`

	// 满足条件的边缘节点个数
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListEdgeNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEdgeNodesResponse struct{}"
	}

	return strings.Join([]string{"ListEdgeNodesResponse", string(data)}, " ")
}
