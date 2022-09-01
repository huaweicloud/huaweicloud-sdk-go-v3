package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListDomainBandwidthPeakResponse struct {

	// 域名对应的带宽峰值列表。
	BandwidthList *[]PeakBandwidthData `json:"bandwidth_list,omitempty" xml:"bandwidth_list"`

	XRequestId     *string `json:"X-Request-Id,omitempty" xml:"X-Request-Id"`
	HttpStatusCode int     `json:"-"`
}

func (o ListDomainBandwidthPeakResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainBandwidthPeakResponse struct{}"
	}

	return strings.Join([]string{"ListDomainBandwidthPeakResponse", string(data)}, " ")
}
