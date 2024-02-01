package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGcbResourceTagsResponse Response Object
type ListGcbResourceTagsResponse struct {

	// 标签列表。
	Tags *[]Tag `json:"tags,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListGcbResourceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGcbResourceTagsResponse struct{}"
	}

	return strings.Join([]string{"ListGcbResourceTagsResponse", string(data)}, " ")
}
