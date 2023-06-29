package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRocketmqTagsResponse Response Object
type ShowRocketmqTagsResponse struct {

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
