package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTagsResponse Response Object
type ListTagsResponse struct {

	// 标签属性
	Tags *[]ResourceTag `json:"tags,omitempty"`

	// 标签数量
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTagsResponse struct{}"
	}

	return strings.Join([]string{"ListTagsResponse", string(data)}, " ")
}
