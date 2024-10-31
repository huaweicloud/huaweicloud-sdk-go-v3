package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DdosAttackLog http攻击日志详情
type DdosAttackLog struct {

	// ddos攻击时间
	AttackTime *int64 `json:"attack_time,omitempty"`

	// 攻击流量带宽平均值
	AvgBps *int64 `json:"avg_bps,omitempty"`

	// 攻击流量带宽峰值
	AvgPps *int64 `json:"avg_pps,omitempty"`

	// 包转发率平均值
	MaxBps *int64 `json:"max_bps,omitempty"`

	// 包转发率峰值
	MaxPps *int64 `json:"max_pps,omitempty"`
}

func (o DdosAttackLog) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DdosAttackLog struct{}"
	}

	return strings.Join([]string{"DdosAttackLog", string(data)}, " ")
}
