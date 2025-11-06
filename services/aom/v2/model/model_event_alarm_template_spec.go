package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EventAlarmTemplateSpec struct {

	// 告警规则类别。
	AlarmSubtype *string `json:"alarm_subtype,omitempty"`

	// 告警规则来源云服务：CCE 创建标识。
	AlarmSource *string `json:"alarm_source,omitempty"`

	// 告警来源。
	EventSource *string `json:"event_source,omitempty"`

	// 监控对象模板（创建告警时需要补齐里面的内容）。
	MonitorObjectTemplates *[]string `json:"monitor_object_templates,omitempty"`

	// 监控对象列表。键值对形式，键值为： - “event_type”：通知类型 - “event_severity”：告警级别 - “event_name”：事件名称 - “namespace”：命名空间 - “clusterId”：集群id - “customField”：用户自定义字段
	MonitorObjects *[]map[string]string `json:"monitor_objects,omitempty"`

	// 触发条件。
	TriggerConditions *[]EventTriggerCondition `json:"trigger_conditions,omitempty"`
}

func (o EventAlarmTemplateSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventAlarmTemplateSpec struct{}"
	}

	return strings.Join([]string{"EventAlarmTemplateSpec", string(data)}, " ")
}
