package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDDoSPeakResponse Response Object
type ShowDDoSPeakResponse struct {

	// 攻击流量
	AttackKbytesTotal *int32 `json:"attack_kbytes_total,omitempty"`

	// 攻击流量峰值包速率
	AttackPpsPeak *int32 `json:"attack_pps_peak,omitempty"`

	// 入流量峰值包速率
	InPpsPeak *int32 `json:"in_pps_peak,omitempty"`

	// 攻击流量峰值
	AttackBpsPeak *int32 `json:"attack_bps_peak,omitempty"`

	// 入流量峰值
	InBpsPeak *int32 `json:"in_bps_peak,omitempty"`

	// 攻击数量
	DdosCount *int32 `json:"ddos_count,omitempty"`

	// 峰值时间
	Utime          *int32 `json:"utime,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowDDoSPeakResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDDoSPeakResponse struct{}"
	}

	return strings.Join([]string{"ShowDDoSPeakResponse", string(data)}, " ")
}
