package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListTagResourcesResponse struct {

	// 标签列表。
	Tags *[]TagDto `json:"tags,omitempty"`

	PageInfo       *PageInfoDto `json:"page_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListTagResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTagResourcesResponse struct{}"
	}

	return strings.Join([]string{"ListTagResourcesResponse", string(data)}, " ")
}
