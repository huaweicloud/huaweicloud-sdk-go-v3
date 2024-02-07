package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInternetBandwidthDomainTagsResponse Response Object
type ListInternetBandwidthDomainTagsResponse struct {

	// 所有标签。
	Tags *[]CreateTag `json:"tags,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListInternetBandwidthDomainTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInternetBandwidthDomainTagsResponse struct{}"
	}

	return strings.Join([]string{"ListInternetBandwidthDomainTagsResponse", string(data)}, " ")
}
