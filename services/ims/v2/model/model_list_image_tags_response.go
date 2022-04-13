package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListImageTagsResponse struct {
	// 标签列表

	Tags           *[]ResourceTag `json:"tags,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListImageTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageTagsResponse struct{}"
	}

	return strings.Join([]string{"ListImageTagsResponse", string(data)}, " ")
}
