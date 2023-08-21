package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStatisticsTrafficResponse Response Object
type ShowStatisticsTrafficResponse struct {

	// 流量数据值，type=max_flow_bandwidth/max_drop_bandwidth/ddos_flow时返回，单位：Kbps
	Value *int64 `json:"value,omitempty"`

	// 入流量，type=flow_drop_traffic时返回
	Flow *[]TimeSeriesData `json:"flow,omitempty"`

	// 清洗流量，type=flow_drop_traffic时返回
	Drop *[]TimeSeriesData `json:"drop,omitempty"`

	// WAF攻击流量，type=attack_traffic时返回
	Waf *[]TimeSeriesData `json:"waf,omitempty"`

	// BOT攻击流量，type=attack_traffic时返回
	Bot *[]TimeSeriesData `json:"bot,omitempty"`

	// CC攻击流量，type=attack_traffic时返回
	Cc *[]TimeSeriesData `json:"cc,omitempty"`

	// DDoS攻击流量，type=attack_traffic时返回
	Ddos           *[]TimeSeriesData `json:"ddos,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowStatisticsTrafficResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStatisticsTrafficResponse struct{}"
	}

	return strings.Join([]string{"ShowStatisticsTrafficResponse", string(data)}, " ")
}
