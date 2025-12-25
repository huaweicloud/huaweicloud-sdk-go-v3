package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlertRuleTemplate struct {

	// 累计次数
	AccumulatedTimes *int32 `json:"accumulated_times,omitempty"`

	// 告警描述
	AlertDescription *string `json:"alert_description,omitempty"`

	// 告警名称
	AlertName *string `json:"alert_name,omitempty"`

	// 告警修复
	AlertRemediation *string `json:"alert_remediation,omitempty"`

	// Map<String,String>
	AlertType map[string]string `json:"alert_type,omitempty"`

	// Iam用户ID
	CreateBy *string `json:"create_by,omitempty"`

	// 毫秒时间戳
	CreateTime *int64 `json:"create_time,omitempty"`

	// Map<String,String>
	CustomProperties map[string]string `json:"custom_properties,omitempty"`

	// 告警规则模板描述
	Description *string `json:"description,omitempty"`

	// 告警组
	EventGrouping *bool `json:"event_grouping,omitempty"`

	JobMode *JobMode `json:"job_mode,omitempty"`

	ProcessStatus *ProcessStatus `json:"process_status,omitempty"`

	// 查询语句
	Query *string `json:"query,omitempty"`

	QueryType *QueryType `json:"query_type,omitempty"`

	Schedule *AlertRuleSchedule `json:"schedule,omitempty"`

	Severity *Serverity `json:"severity,omitempty"`

	// 是否仿真
	Simulation *bool `json:"simulation,omitempty"`

	Status *Status `json:"status,omitempty"`

	// 告警抑制
	Suppresion *bool `json:"suppresion,omitempty"`

	// 表名称
	TableName *string `json:"table_name,omitempty"`

	// UUID
	TemplateId *string `json:"template_id,omitempty"`

	// 模板名称
	TemplateName *string `json:"template_name,omitempty"`

	// 触发器数组
	Triggers *[]Trigger `json:"triggers,omitempty"`

	// Iam用户ID
	UpdateBy *string `json:"update_by,omitempty"`

	// 毫秒时间戳
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 毫秒时间戳
	UpdateTimeByUser *int64 `json:"update_time_by_user,omitempty"`
}

func (o AlertRuleTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlertRuleTemplate struct{}"
	}

	return strings.Join([]string{"AlertRuleTemplate", string(data)}, " ")
}
