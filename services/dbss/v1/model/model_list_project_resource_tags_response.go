package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectResourceTagsResponse Response Object
type ListProjectResourceTagsResponse struct {

	// 标签列表
	Tags           *[]ProjectResourceTagResponseTags `json:"tags,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o ListProjectResourceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectResourceTagsResponse struct{}"
	}

	return strings.Join([]string{"ListProjectResourceTagsResponse", string(data)}, " ")
}
