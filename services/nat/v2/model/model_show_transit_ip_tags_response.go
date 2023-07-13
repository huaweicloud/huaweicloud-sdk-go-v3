package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTransitIpTagsResponse Response Object
type ShowTransitIpTagsResponse struct {

	// 请求id。
	RequestId *string `json:"request_id,omitempty"`

	// 标签。
	Tags           *[]Tag `json:"tags,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowTransitIpTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTransitIpTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowTransitIpTagsResponse", string(data)}, " ")
}
