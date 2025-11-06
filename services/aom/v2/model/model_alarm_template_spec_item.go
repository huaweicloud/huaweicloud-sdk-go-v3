package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlarmTemplateSpecItem struct {

	// 告警模板下单个告警规则名称。
	AlarmTemplateName string `json:"alarm_template_name"`

	// 告警模板下单个告警规则英文名称。
	AlarmTemplateNameEn *string `json:"alarm_template_name_en,omitempty"`

	// 告警模板下单个告警规则描述。
	Desc *string `json:"desc,omitempty"`

	// 告警模板下单个告警规则英文描述。
	DescEn *string `json:"desc_en,omitempty"`

	// 告警模板下单个告警规则类型。 “metric”：指标告警 “event”：事件告警
	AlarmTemplateSpecType string `json:"alarm_template_spec_type"`

	MetricAlarmTemplateSpec *MetricAlarmTemplateSpec `json:"metric_alarm_template_spec,omitempty"`

	EventAlarmTemplateSpec *EventAlarmTemplateSpec `json:"event_alarm_template_spec,omitempty"`
}

func (o AlarmTemplateSpecItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmTemplateSpecItem struct{}"
	}

	return strings.Join([]string{"AlarmTemplateSpecItem", string(data)}, " ")
}
