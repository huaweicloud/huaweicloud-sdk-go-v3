package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EIP防护流量响应体
type DailyData struct {

	// 开始时间
	PeriodStart int64 `json:"period_start" xml:"period_start"`

	// 入流量（bit/s）
	BpsIn int32 `json:"bps_in" xml:"bps_in"`

	// 攻击流量（bit/s）
	BpsAttack int64 `json:"bps_attack" xml:"bps_attack"`

	// 总流量
	TotalBps int64 `json:"total_bps" xml:"total_bps"`

	// 入报文速率（个/s）
	PpsIn int64 `json:"pps_in" xml:"pps_in"`

	// 攻击文速率（个/s）
	PpsAttack int64 `json:"pps_attack" xml:"pps_attack"`

	// 总报文速率
	TotalPps int64 `json:"total_pps" xml:"total_pps"`
}

func (o DailyData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DailyData struct{}"
	}

	return strings.Join([]string{"DailyData", string(data)}, " ")
}
