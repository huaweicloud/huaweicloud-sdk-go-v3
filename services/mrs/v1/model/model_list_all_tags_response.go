package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAllTagsResponse struct {
	// 标签列表信息

	Tags           *[]TagWithMultiValue `json:"tags,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListAllTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllTagsResponse struct{}"
	}

	return strings.Join([]string{"ListAllTagsResponse", string(data)}, " ")
}
