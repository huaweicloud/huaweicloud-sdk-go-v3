package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRocketmqTagsResponse Response Object
type ShowRocketmqTagsResponse struct {

	// 总数。
	Total float32 `json:"total,omitempty"`

	// 下个分页的offset。
	NextOffset *int32 `json:"next_offset,omitempty"`

	// 上个分页的offset。
	PreviousOffset *int32 `json:"previous_offset,omitempty"`

	// 标签列表
	Tags           *[]TagEntity `json:"tags,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowRocketmqTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRocketmqTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowRocketmqTagsResponse", string(data)}, " ")
}
