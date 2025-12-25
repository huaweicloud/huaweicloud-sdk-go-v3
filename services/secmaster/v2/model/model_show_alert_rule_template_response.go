package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAlertRuleTemplateResponse Response Object
type ShowAlertRuleTemplateResponse struct {

	// 累计次数
	AccumulatedTimes *int32 `json:"accumulated_times,omitempty"`

	// UUID
	CreateBy *string `json:"create_by,omitempty"`

	// 毫秒时间戳
	CreateTime *int64 `json:"create_time,omitempty"`

	// 数量
	CuQuotaAmount float32 `json:"cu_quota_amount,omitempty"`

	// 告警规则模板描述
	Description *string `json:"description,omitempty"`

	Environment *AlertRuleEnvironment `json:"environment,omitempty"`

	JobMode *IsapJobMode `json:"job_mode,omitempty"`

	JobModeSetting *IsapJobModeSettingVo `json:"job_mode_setting,omitempty"`

	JobOutputSetting *IsapJobOutputSetting `json:"job_output_setting,omitempty"`

	ProcessError *ProcessError `json:"process_error,omitempty"`

	ProcessStatus *ProcessStatus `json:"process_status,omitempty"`

	QueryType *QueryType `json:"query_type,omitempty"`

	// Script 脚本
	Script *string `json:"script,omitempty"`

	Status *Status `json:"status,omitempty"`

	// 表名称
	TableName *string `json:"table_name,omitempty"`

	// UUID
	TemplateId *string `json:"template_id,omitempty"`

	// 模板名称
	TemplateName *string `json:"template_name,omitempty"`

	// 触发器数组
	Triggers *[]Trigger `json:"triggers,omitempty"`

	// UUID
	UpdateBy *string `json:"update_by,omitempty"`

	// 毫秒时间戳
	UpdateTime     *int64 `json:"update_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowAlertRuleTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlertRuleTemplateResponse struct{}"
	}

	return strings.Join([]string{"ShowAlertRuleTemplateResponse", string(data)}, " ")
}
