package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDomainBandwidthPeakResponse Response Object
type ListDomainBandwidthPeakResponse struct {

	// 域名对应的带宽峰值列表。
	BandwidthList *[]PeakBandwidthData `json:"bandwidth_list,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListDomainBandwidthPeakResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainBandwidthPeakResponse struct{}"
	}

	return strings.Join([]string{"ListDomainBandwidthPeakResponse", string(data)}, " ")
}
