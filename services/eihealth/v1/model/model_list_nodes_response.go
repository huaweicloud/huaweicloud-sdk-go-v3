package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNodesResponse Response Object
type ListNodesResponse struct {

	// 节点总数
	Count *int32 `json:"count,omitempty"`

	// 节点列表
	Nodes          *[]string `json:"nodes,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNodesResponse struct{}"
	}

	return strings.Join([]string{"ListNodesResponse", string(data)}, " ")
}
