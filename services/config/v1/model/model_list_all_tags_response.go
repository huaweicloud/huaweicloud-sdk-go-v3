package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAllTagsResponse Response Object
type ListAllTagsResponse struct {

	// 标签列表
	Tags *[]TagDetail `json:"tags,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListAllTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllTagsResponse struct{}"
	}

	return strings.Join([]string{"ListAllTagsResponse", string(data)}, " ")
}
