package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceShareTagsResponse Response Object
type ListResourceShareTagsResponse struct {

	// 一个或多个标签键值对的列表。标签键必须存在，而不是空字符串。标签值必须存在，但可以是空字符串。
	Tags *[]TagDto `json:"tags,omitempty"`

	PageInfo       *PageInfoMarkerByKey `json:"page_info,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListResourceShareTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceShareTagsResponse struct{}"
	}

	return strings.Join([]string{"ListResourceShareTagsResponse", string(data)}, " ")
}
