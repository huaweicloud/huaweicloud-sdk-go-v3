package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourcesByTagsRequestBody 使用标签过滤实例请求体。
type ListResourcesByTagsRequestBody struct {

	// 是否不包含任意一个标签
	WithoutAnyTag *bool `json:"without_any_tag,omitempty"`

	// 标签列表
	Tags *[]TagWithValues `json:"tags,omitempty"`

	// 搜索字段列表
	Matches *[]Match `json:"matches,omitempty"`
}

func (o ListResourcesByTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourcesByTagsRequestBody struct{}"
	}

	return strings.Join([]string{"ListResourcesByTagsRequestBody", string(data)}, " ")
}
