package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTagWorkspacesRequestBody 标签管理界面查询资源实例列表的请求体。
type ListTagWorkspacesRequestBody struct {

	// 不包含任意一个标签，该字段为true时查询所有不带标签的资源，此时忽略 “tags”、“tags_any”、“not_tags”、“not_tags_any”字段
	WithoutAnyTag *bool `json:"without_any_tag,omitempty"`

	// 包含标签，最多包含50个key，每个key下面的value最多10个，每个key对应的value可以为空数组但结构体不能缺失。
	Tags *[]ResourceTagParam `json:"tags,omitempty"`

	// 搜索字段，key为要匹配的字段，如resource_name等。value为匹配的值。
	Matches *[]TagMatch `json:"matches,omitempty"`

	// 系统标签列表，目前只包含一个tag结构体。key下面只包含一个value。
	SysTags *[]SystemTagParam `json:"sys_tags,omitempty"`
}

func (o ListTagWorkspacesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTagWorkspacesRequestBody struct{}"
	}

	return strings.Join([]string{"ListTagWorkspacesRequestBody", string(data)}, " ")
}
