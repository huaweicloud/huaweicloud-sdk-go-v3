package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNodesResponse Response Object
type ListNodesResponse struct {

	// DDM实例节点信息列表的集合。
	Nodes *[]NodeList `json:"nodes,omitempty"`

	// 分页参数: 起始值。
	Offset *int32 `json:"offset,omitempty"`

	// 分页参数：每页多少条。
	Limit *int32 `json:"limit,omitempty"`

	// DDM实例节点个数。
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNodesResponse struct{}"
	}

	return strings.Join([]string{"ListNodesResponse", string(data)}, " ")
}
