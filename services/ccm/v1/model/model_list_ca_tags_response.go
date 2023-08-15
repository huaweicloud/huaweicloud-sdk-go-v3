package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCaTagsResponse Response Object
type ListCaTagsResponse struct {

	// 标签列表，key和value键值对的集合。
	Tags           *[]ResourceTag `json:"tags,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListCaTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCaTagsResponse struct{}"
	}

	return strings.Join([]string{"ListCaTagsResponse", string(data)}, " ")
}
