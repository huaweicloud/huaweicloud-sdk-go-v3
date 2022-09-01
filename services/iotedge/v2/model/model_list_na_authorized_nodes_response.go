package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListNaAuthorizedNodesResponse struct {

	// 总记录数
	Count *int64 `json:"count,omitempty" xml:"count"`

	PageInfo *PageInfoDto `json:"page_info,omitempty" xml:"page_info"`

	// 查询授权北向NA信息到边缘节点列表的返回结构体
	Nodes          *[]QueryAuthorizedNodeDto `json:"nodes,omitempty" xml:"nodes"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListNaAuthorizedNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNaAuthorizedNodesResponse struct{}"
	}

	return strings.Join([]string{"ListNaAuthorizedNodesResponse", string(data)}, " ")
}
