package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MetricAlarmTemplateSpec struct {

	// 告警规则类别。
	AlarmSubtype *string `json:"alarm_subtype,omitempty"`

	// 告警规则来源云服务：CCE 创建标识。
	AlarmSource *string `json:"alarm_source,omitempty"`

	// 监控类型。
	MonitorType *string `json:"monitor_type,omitempty"`

	// 触发条件。
	TriggerConditions *[]TemplateTriggerCondition `json:"trigger_conditions,omitempty"`

	// 数据不足条件。
	NoDataConditions *[]NoDataCondition `json:"no_data_conditions,omitempty"`

	// 告警标签。
	AlarmTags *[]AlarmTags `json:"alarm_tags,omitempty"`

	RecoveryConditions *RecoveryCondition `json:"recovery_conditions,omitempty"`
}

func (o MetricAlarmTemplateSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricAlarmTemplateSpec struct{}"
	}

	return strings.Join([]string{"MetricAlarmTemplateSpec", string(data)}, " ")
}
