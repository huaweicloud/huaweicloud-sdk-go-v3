package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlarmLogResponseAlarmLog struct {

	// 告警ID
	Id *string `json:"id,omitempty"`

	// 告警状态 - ON - OFF
	AlarmLife *string `json:"alarmLife,omitempty"`

	// 是否发送邮件
	SendEmail *bool `json:"sendEmail,omitempty"`

	// 告警发生时间
	AlarmTime *string `json:"alarm_time,omitempty"`

	// 告警类型 - RISK_RULE: 风险规则 - RISK_CPU: CPU超限 - RISK_MEMORY: 内存超限 - RISK_DISK: 磁盘超限 - RISK_DISK_CAPACITY: 磁盘容量不足六个月 - RISK_BACKUP: 备份失败 - AUDIT_QPS_OVERFLOW: 流量超限入库延迟告警 - RISK_AGENT: Agent异常 - AUDIT_BACKUP_FAILED: 实例备份失败(运维侧)
	AlarmType *string `json:"alarm_type,omitempty"`

	// 告警恢复时间
	AlarmFixTime *string `json:"alarm_fix_time,omitempty"`

	// 告警确认状态 - DONE: 已确认 - UNDO: 未确认
	AlarmStatus *string `json:"alarm_status,omitempty"`

	// 告警风险等级 - LOW - MEDIUM - HIGH
	AlarmRisk *string `json:"alarm_risk,omitempty"`

	// 告警描述信息
	AlarmDescription *string `json:"alarm_description,omitempty"`
}

func (o AlarmLogResponseAlarmLog) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmLogResponseAlarmLog struct{}"
	}

	return strings.Join([]string{"AlarmLogResponseAlarmLog", string(data)}, " ")
}
