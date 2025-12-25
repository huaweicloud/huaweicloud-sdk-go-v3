package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlertRuleJobSetting 告警规则作业设置
type AlertRuleJobSetting struct {

	// 映射表
	AlertCustomProperties map[string]string `json:"alert_custom_properties,omitempty"`

	// 告警描述
	AlertDescription *string `json:"alert_description,omitempty"`

	// 分组标志
	AlertGrouping *bool `json:"alert_grouping,omitempty"`

	// 映射表
	AlertMapping map[string]string `json:"alert_mapping,omitempty"`

	// 告警报名称
	AlertName *string `json:"alert_name,omitempty"`

	// 告警修复建议
	AlertRemediation *string `json:"alert_remediation,omitempty"`

	AlertSeverity *IsapAlertSeverity `json:"alert_severity,omitempty"`

	// 抑制标志
	AlertSuppression *bool `json:"alert_suppression,omitempty"`

	// 告警类型映射表
	AlertType map[string]string `json:"alert_type,omitempty"`

	// 提取的实体
	EntityExtraction map[string]string `json:"entity_extraction,omitempty"`

	// 字段映射
	FieldMapping map[string]string `json:"field_mapping,omitempty"`
}

func (o AlertRuleJobSetting) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlertRuleJobSetting struct{}"
	}

	return strings.Join([]string{"AlertRuleJobSetting", string(data)}, " ")
}
