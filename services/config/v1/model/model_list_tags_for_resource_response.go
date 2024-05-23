package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTagsForResourceResponse Response Object
type ListTagsForResourceResponse struct {

	// 标签列表
	Tags           *[]ResourceTag `json:"tags,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListTagsForResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTagsForResourceResponse struct{}"
	}

	return strings.Join([]string{"ListTagsForResourceResponse", string(data)}, " ")
}
