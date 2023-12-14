package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTrackedResourceTagsResponse Response Object
type ListTrackedResourceTagsResponse struct {

	// 标签列表
	Tags *[]TagDetail `json:"tags,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListTrackedResourceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTrackedResourceTagsResponse struct{}"
	}

	return strings.Join([]string{"ListTrackedResourceTagsResponse", string(data)}, " ")
}
