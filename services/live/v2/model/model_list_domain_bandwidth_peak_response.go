package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListDomainBandwidthPeakResponse struct {
	// 域名对应的带宽峰值列表。

	BandwidthList *[]PeakBandwidthData `json:"bandwidth_list,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListDomainBandwidthPeakResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainBandwidthPeakResponse struct{}"
	}

	return strings.Join([]string{"ListDomainBandwidthPeakResponse", string(data)}, " ")
}
