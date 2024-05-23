package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTagsForResourceTypeResponse Response Object
type ListTagsForResourceTypeResponse struct {

	// 标签列表
	Tags *[]Tag `json:"tags,omitempty"`

	// 总记录数
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListTagsForResourceTypeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTagsForResourceTypeResponse struct{}"
	}

	return strings.Join([]string{"ListTagsForResourceTypeResponse", string(data)}, " ")
}
