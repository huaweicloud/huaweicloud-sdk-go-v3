package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListProductTopicsResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty" xml:"total"`

	// 本次返回数量
	Size *int32 `json:"size,omitempty" xml:"size"`

	// 产品主题列表
	Items          *[]ProductTopic `json:"items,omitempty" xml:"items"`
	HttpStatusCode int             `json:"-"`
}

func (o ListProductTopicsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProductTopicsResponse struct{}"
	}

	return strings.Join([]string{"ListProductTopicsResponse", string(data)}, " ")
}
