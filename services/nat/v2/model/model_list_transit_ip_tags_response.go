package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListTransitIpTagsResponse struct {

	// 请求id。
	RequestId *string `json:"request_id,omitempty"`

	// 标签。
	Tags           *[]Tags `json:"tags,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListTransitIpTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTransitIpTagsResponse struct{}"
	}

	return strings.Join([]string{"ListTransitIpTagsResponse", string(data)}, " ")
}
