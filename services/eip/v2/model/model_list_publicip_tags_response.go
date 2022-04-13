package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListPublicipTagsResponse struct {
	// 标签列表

	Tags           *[]TagResp `json:"tags,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListPublicipTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPublicipTagsResponse struct{}"
	}

	return strings.Join([]string{"ListPublicipTagsResponse", string(data)}, " ")
}
