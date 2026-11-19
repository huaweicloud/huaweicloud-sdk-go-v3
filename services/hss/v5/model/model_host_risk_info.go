package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HostRiskInfo struct {

	// **参数解释**: 基线风险数量 **取值范围**: 最小值0，最大值2147483647
	BaselineRiskNum *int32 `json:"baseline_risk_num,omitempty"`

	// **参数解释**: 漏洞总数 **取值范围**: 最小值0，最大值2147483647
	VulTotalNum *int32 `json:"vul_total_num,omitempty"`

	// **参数解释**: 紧急漏洞数量 **取值范围**: 最小值0，最大值2147483647
	CriticalVulNum *int32 `json:"critical_vul_num,omitempty"`

	// **参数解释**: 告警总数 **取值范围**: 最小值0，最大值2147483647
	AlarmTotalNum *int32 `json:"alarm_total_num,omitempty"`

	// **参数解释**: 需紧急处理的告警（危急+高危） **取值范围**: 最小值0，最大值2147483647
	CriticalAlarmNum *int32 `json:"critical_alarm_num,omitempty"`

	// **参数解释**: 待处理可疑进程数量 **取值范围**: 最小值0，最大值2147483647
	UnhandledSuspiciousProcessNum *int32 `json:"unhandled_suspicious_process_num,omitempty"`
}

func (o HostRiskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostRiskInfo struct{}"
	}

	return strings.Join([]string{"HostRiskInfo", string(data)}, " ")
}
