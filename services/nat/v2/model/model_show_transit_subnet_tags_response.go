package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTransitSubnetTagsResponse Response Object
type ShowTransitSubnetTagsResponse struct {

	// 请求id。
	RequestId *string `json:"request_id,omitempty"`

	// 标签。
	Tags           *[]Tag `json:"tags,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowTransitSubnetTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTransitSubnetTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowTransitSubnetTagsResponse", string(data)}, " ")
}
