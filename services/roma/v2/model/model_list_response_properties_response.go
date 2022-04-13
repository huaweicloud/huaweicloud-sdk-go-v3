package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListResponsePropertiesResponse struct {
	// 总数

	Total *int32 `json:"total,omitempty"`
	// 本次返回数量

	Size *int32 `json:"size,omitempty"`
	// 属性列表

	Items          *[]Property `json:"items,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListResponsePropertiesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResponsePropertiesResponse struct{}"
	}

	return strings.Join([]string{"ListResponsePropertiesResponse", string(data)}, " ")
}
