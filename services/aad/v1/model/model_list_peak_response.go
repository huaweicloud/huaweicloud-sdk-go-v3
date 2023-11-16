package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPeakResponse Response Object
type ListPeakResponse struct {

	// 攻击峰值
	AttackKbpsPeak float32 `json:"attack_kbps_peak,omitempty"`

	// 流量峰值
	InKbpsPeak float32 `json:"in_kbps_peak,omitempty"`

	// 攻击次数
	DdosCount float32 `json:"ddos_count,omitempty"`

	// 攻击峰值发生时间点
	Timestamp float32 `json:"timestamp,omitempty"`

	// 高防IP
	Vip            *string `json:"vip,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListPeakResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPeakResponse struct{}"
	}

	return strings.Join([]string{"ListPeakResponse", string(data)}, " ")
}
