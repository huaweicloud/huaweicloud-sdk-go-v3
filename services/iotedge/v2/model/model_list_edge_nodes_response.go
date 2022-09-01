package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListEdgeNodesResponse struct {

	// 总记录数
	Count *int64 `json:"count,omitempty" xml:"count"`

	PageInfo *PageInfoDto `json:"page_info,omitempty" xml:"page_info"`

	// 节点列表
	Nodes          *[]EdgeNodeDto `json:"nodes,omitempty" xml:"nodes"`
	HttpStatusCode int            `json:"-"`
}

func (o ListEdgeNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEdgeNodesResponse struct{}"
	}

	return strings.Join([]string{"ListEdgeNodesResponse", string(data)}, " ")
}
