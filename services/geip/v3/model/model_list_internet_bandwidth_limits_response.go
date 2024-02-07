package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInternetBandwidthLimitsResponse Response Object
type ListInternetBandwidthLimitsResponse struct {

	// 本次请求的编号
	RequestId *string `json:"request_id,omitempty"`

	// 全域公网带宽限制列表
	InternetBandwidthLimits *[]ListInternetBandwidthLimits `json:"internet_bandwidth_limits,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListInternetBandwidthLimitsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInternetBandwidthLimitsResponse struct{}"
	}

	return strings.Join([]string{"ListInternetBandwidthLimitsResponse", string(data)}, " ")
}
