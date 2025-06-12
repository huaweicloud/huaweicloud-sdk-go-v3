package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourcesTagsResponse Response Object
type ListResourcesTagsResponse struct {

	// 标签列表
	Tags           *[]Tag `json:"tags,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListResourcesTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourcesTagsResponse struct{}"
	}

	return strings.Join([]string{"ListResourcesTagsResponse", string(data)}, " ")
}
