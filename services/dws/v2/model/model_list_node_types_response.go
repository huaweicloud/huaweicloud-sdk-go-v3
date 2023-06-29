package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNodeTypesResponse Response Object
type ListNodeTypesResponse struct {

	// 节点类型对象列表。
	NodeTypes *[]NodeTypes `json:"node_types,omitempty"`

	// 节点类型总数
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListNodeTypesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNodeTypesResponse struct{}"
	}

	return strings.Join([]string{"ListNodeTypesResponse", string(data)}, " ")
}
