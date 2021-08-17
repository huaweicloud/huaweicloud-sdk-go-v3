package model

import (
	"encoding/json"

	"strings"
)

// WeeklyCount
type WeeklyCount struct {
	// DDoS拦截次数

	DdosInterceptTimes int32 `json:"ddos_intercept_times"`
	// DDoS黑洞次数

	DdosBlackholeTimes int32 `json:"ddos_blackhole_times"`
	// 最大攻击流量

	MaxAttackBps int32 `json:"max_attack_bps"`
	// 最大攻击连接数

	MaxAttackConns int32 `json:"max_attack_conns"`
	// 开始时间

	PeriodStartDate int64 `json:"period_start_date"`
}

func (o WeeklyCount) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "WeeklyCount struct{}"
	}

	return strings.Join([]string{"WeeklyCount", string(data)}, " ")
}
