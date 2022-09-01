package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListTagValuesResponse struct {

	// 标签的名称
	TagName *string `json:"tag_name,omitempty" xml:"tag_name"`

	// 标签的值列表
	TagValues *[]string `json:"tag_values,omitempty" xml:"tag_values"`

	// 当前列表元素总数
	Count          *int32 `json:"count,omitempty" xml:"count"`
	HttpStatusCode int    `json:"-"`
}

func (o ListTagValuesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTagValuesResponse struct{}"
	}

	return strings.Join([]string{"ListTagValuesResponse", string(data)}, " ")
}
