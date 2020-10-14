/*
 * LiveAPI
 *
 * 直播服务源站所有接口
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListDomainBandwidthSummaryV2Response struct {
	// 域名对应的带宽峰值列表。
	BandwidthList *[]PeakBandwidthData `json:"bandwidth_list,omitempty"`
	XRequestId    *string              `json:"X-request-id,omitempty"`
}

func (o ListDomainBandwidthSummaryV2Response) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListDomainBandwidthSummaryV2Response", string(data)}, " ")
}
