package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRequestPropertiesResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty" xml:"total"`

	// 本次返回数量
	Size *int32 `json:"size,omitempty" xml:"size"`

	// 属性列表
	Items          *[]Property `json:"items,omitempty" xml:"items"`
	HttpStatusCode int         `json:"-"`
}

func (o ListRequestPropertiesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRequestPropertiesResponse struct{}"
	}

	return strings.Join([]string{"ListRequestPropertiesResponse", string(data)}, " ")
}
