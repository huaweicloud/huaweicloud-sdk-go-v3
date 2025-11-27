package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTransitSubnetTagsResponse Response Object
type ListTransitSubnetTagsResponse struct {

	// 请求id。
	RequestId *string `json:"request_id,omitempty"`

	// 标签。
	Tags           *[]Tags `json:"tags,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListTransitSubnetTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTransitSubnetTagsResponse struct{}"
	}

	return strings.Join([]string{"ListTransitSubnetTagsResponse", string(data)}, " ")
}
