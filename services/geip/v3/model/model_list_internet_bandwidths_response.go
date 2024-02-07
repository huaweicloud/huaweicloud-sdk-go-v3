package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInternetBandwidthsResponse Response Object
type ListInternetBandwidthsResponse struct {

	// 本次请求的编号
	RequestId *string `json:"request_id,omitempty"`

	// 全域公网带宽列表
	InternetBandwidths *[]ListInternetBandwidths `json:"internet_bandwidths,omitempty"`

	PageInfo *ListGlobalEipsResponseBodyPageInfo `json:"page_info,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListInternetBandwidthsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInternetBandwidthsResponse struct{}"
	}

	return strings.Join([]string{"ListInternetBandwidthsResponse", string(data)}, " ")
}
