package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OperationSummaryInfo struct {

	// **参数解释**: 用户访问HSS天数 **取值范围**: 最小值0，最大值365
	HssVisitDays *int32 `json:"hss_visit_days,omitempty"`

	// **参数解释**: 某个工作超过的用户百分比（目前只有漏洞或告警，哪个打败的用户占比高，显示哪个） **取值范围**: 最小值0，最大值1
	WorkloadBeatRate *float32 `json:"workload_beat_rate,omitempty"`

	// **参数解释**: 用户名 **取值范围**: 字符长度0-128位
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**: 当前月初时间戳 **取值范围**: 最小值0，最大值9223372036854775807
	CurrentMonthStart *int64 `json:"current_month_start,omitempty"`

	// **参数解释**: 当前月末时间戳 **取值范围**: 最小值0，最大值9223372036854775807
	CurrentMonthEnd *int64 `json:"current_month_end,omitempty"`

	// **参数解释**: 处置的安全事件条数 **取值范围**: 最小值0，最大值2147483647
	HandledSecurityEventNum *int64 `json:"handled_security_event_num,omitempty"`

	// **参数解释**: 工作量打败的用户百分比 **取值范围**: 最小值0，最大值1
	TotalWorkloadBeatRate *float32 `json:"total_workload_beat_rate,omitempty"`

	// **参数解释**:  称号 **取值范围**: -vul-fix-master: 补洞大师。 -vul-fix-expert: 漏洞修复小能手。 -baseline-handle: 风险配置处置高手。 -malware-file: 防病毒先锋。 -ransomware-event: 防勒索达人。 -web-tamper-event: 网站守卫。
	Title *string `json:"title,omitempty"`

	// **参数解释**: 时间字符串 **取值范围**: 字符长度0-32位
	ReportId *string `json:"report_id,omitempty"`

	// **参数解释**: 当前月份 **取值范围**: 最小值1，最大值12
	CurrentMonth *int32 `json:"current_month,omitempty"`

	// **参数解释**: 工作（哪个打败的用户占比高，显示哪个) **取值范围**: -vul-fix: 漏洞修复 -baseline-handle: 基线处置 -event-handle: 处置入侵事件
	Work *string `json:"work,omitempty"`

	// **参数解释**: 报告生成时间戳 **取值范围**: 最小值0，最大值9223372036854775807
	CreateTime *int64 `json:"create_time,omitempty"`
}

func (o OperationSummaryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperationSummaryInfo struct{}"
	}

	return strings.Join([]string{"OperationSummaryInfo", string(data)}, " ")
}
