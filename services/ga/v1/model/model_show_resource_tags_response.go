package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowResourceTagsResponse struct {

	// 标签列表。
	Tags *[]ResourceTag `json:"tags,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowResourceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowResourceTagsResponse", string(data)}, " ")
}
