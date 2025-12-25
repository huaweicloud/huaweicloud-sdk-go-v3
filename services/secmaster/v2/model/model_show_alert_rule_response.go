package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAlertRuleResponse Response Object
type ShowAlertRuleResponse struct {

	// UUID
	AlertRuleId *string `json:"alert_rule_id,omitempty"`

	// Alert rule name 告警规则名称
	AlertRuleName *string `json:"alert_rule_name,omitempty"`

	// Job Script 作业脚本
	Script *string `json:"script,omitempty"`

	Status *JobStatus `json:"status,omitempty"`

	// directory 目录分组
	Directory *string `json:"directory,omitempty"`

	// Alert rule description 告警规则描述
	Description *string `json:"description,omitempty"`

	JobMode *IsapJobMode `json:"job_mode,omitempty"`

	JobModeSetting *IsapJobModeSettingVo `json:"job_mode_setting,omitempty"`

	JobOutputSetting *AlertRuleJobSetting `json:"job_output_setting,omitempty"`

	ProcessStatus *JobProcessStatus `json:"process_status,omitempty"`

	ProcessError *AlertRuleProcessError `json:"process_error,omitempty"`

	Environment *JobEnvironment `json:"environment,omitempty"`

	// UUID
	OutputTableId *string `json:"output_table_id,omitempty"`

	// 表名称
	OutputTableName *string `json:"output_table_name,omitempty"`

	// 输出表ID列表
	OutputTableIds *[]string `json:"output_table_ids,omitempty"`

	// 输出表名称列表
	OutputTableNames *[]string `json:"output_table_names,omitempty"`

	// 创建者
	CreateBy *string `json:"create_by,omitempty"`

	// 毫秒时间戳
	CreateTime *int64 `json:"create_time,omitempty"`

	// 更新者
	UpdateBy *string `json:"update_by,omitempty"`

	// 毫秒时间戳
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 毫秒时间戳
	DeleteTime     *int64 `json:"delete_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowAlertRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlertRuleResponse struct{}"
	}

	return strings.Join([]string{"ShowAlertRuleResponse", string(data)}, " ")
}
