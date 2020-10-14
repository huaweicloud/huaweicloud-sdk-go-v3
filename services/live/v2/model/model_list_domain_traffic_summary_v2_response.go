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
type ListDomainTrafficSummaryV2Response struct {
	// 域名对应的流量汇总列表。
	TrafficList *[]TrafficSummaryData `json:"traffic_list,omitempty"`
	XRequestId  *string               `json:"X-request-id,omitempty"`
}

func (o ListDomainTrafficSummaryV2Response) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListDomainTrafficSummaryV2Response", string(data)}, " ")
}
