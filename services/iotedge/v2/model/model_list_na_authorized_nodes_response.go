package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListNaAuthorizedNodesResponse struct {
	// 总记录数

	Count *int64 `json:"count,omitempty"`

	PageInfo *PageInfoDto `json:"page_info,omitempty"`
	// 查询授权北向NA信息到边缘节点列表的返回结构体

	Nodes          *[]QueryAuthorizedNodeDto `json:"nodes,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListNaAuthorizedNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNaAuthorizedNodesResponse struct{}"
	}

	return strings.Join([]string{"ListNaAuthorizedNodesResponse", string(data)}, " ")
}
