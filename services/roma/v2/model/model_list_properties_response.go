package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListPropertiesResponse struct {
	// 总数

	Total *int32 `json:"total,omitempty"`
	// 本次返回数量

	Size *int32 `json:"size,omitempty"`
	// 属性列表

	Items          *[]Property `json:"items,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListPropertiesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPropertiesResponse struct{}"
	}

	return strings.Join([]string{"ListPropertiesResponse", string(data)}, " ")
}
