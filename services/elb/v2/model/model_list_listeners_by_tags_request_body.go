package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type ListListenersByTagsRequestBody struct {
	// 分页起始。

	Offset *int32 `json:"offset,omitempty"`
	// 分页大小。

	Limit *int32 `json:"limit,omitempty"`
	// 操作标识（仅限于filter，count）： filter（过滤），如果是filter就是分页查询 count(查询总条数)，按照条件将总条数返回。

	Action string `json:"action"`
	// 搜索字段，key为要匹配的字段，如resource_name等。value为匹配的值。key为固定字典值。根据不同的字段确认是否需要模糊匹配，如resource_name默认为模糊搜索，如果value为空字符串精确匹配。key如果是resource_id则精确匹配。

	Matches *[]ActionMatch `json:"matches,omitempty"`
	// 要搜索的标签值

	Tags *[]ActionTag `json:"tags,omitempty"`
	// 查询不包含任何标签的资源，该字段为true时，忽略tags字段的查询条件。

	WithoutAnyTag *bool `json:"without_any_tag,omitempty"`
}

func (o ListListenersByTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListListenersByTagsRequestBody struct{}"
	}

	return strings.Join([]string{"ListListenersByTagsRequestBody", string(data)}, " ")
}
