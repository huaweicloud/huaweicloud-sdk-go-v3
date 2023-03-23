package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListPrivateNatTagsResponse struct {

	// 请求id。
	RequestId *string `json:"request_id,omitempty"`

	// 标签。
	Tags           *[]Tags `json:"tags,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListPrivateNatTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPrivateNatTagsResponse struct{}"
	}

	return strings.Join([]string{"ListPrivateNatTagsResponse", string(data)}, " ")
}
