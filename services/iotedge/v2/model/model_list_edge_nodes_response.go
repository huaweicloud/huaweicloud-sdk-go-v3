package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListEdgeNodesResponse struct {
	// 总记录数

	Count *int64 `json:"count,omitempty"`

	PageInfo *PageInfoDto `json:"page_info,omitempty"`
	// 节点列表

	Nodes          *[]EdgeNodeDto `json:"nodes,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListEdgeNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEdgeNodesResponse struct{}"
	}

	return strings.Join([]string{"ListEdgeNodesResponse", string(data)}, " ")
}
