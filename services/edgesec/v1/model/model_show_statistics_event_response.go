package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStatisticsEventResponse Response Object
type ShowStatisticsEventResponse struct {

	// DDos攻击事件次数，type=ddos_attack_count返回
	Value *int64 `json:"value,omitempty"`

	// WAF攻击事件次数，type=attack_count时返回
	Waf *[]TimeSeriesData `json:"waf,omitempty"`

	// BOT攻击事件次数，type=attack_count时返回
	Bot *[]TimeSeriesData `json:"bot,omitempty"`

	// CC攻击事件次数，type=attack_count时返回
	Cc *[]TimeSeriesData `json:"cc,omitempty"`

	// DDos攻击事件次数，type=attack_count返回
	Ddos *[]TimeSeriesData `json:"ddos,omitempty"`

	// 访问次数，type=flow_drop_count返回
	Flow *[]TimeSeriesData `json:"flow,omitempty"`

	// 攻击次数，type=flow_drop_count返回
	Drop           *[]TimeSeriesData `json:"drop,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowStatisticsEventResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStatisticsEventResponse struct{}"
	}

	return strings.Join([]string{"ShowStatisticsEventResponse", string(data)}, " ")
}
