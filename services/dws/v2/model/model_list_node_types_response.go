package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListNodeTypesResponse struct {
	// 节点类型对象列表

	NodeTypes      *[]NodeTypes `json:"node_types,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListNodeTypesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNodeTypesResponse struct{}"
	}

	return strings.Join([]string{"ListNodeTypesResponse", string(data)}, " ")
}
