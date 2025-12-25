package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateAlertRuleRequestBody struct {

	// Alert rule name 告警规则名称
	AlertRuleName *string `json:"alert_rule_name,omitempty"`

	// Alert rule description 告警规则描述
	Description *string `json:"description,omitempty"`

	// directory 目录分组
	Directory *string `json:"directory,omitempty"`

	// Job Script 作业脚本
	Script *string `json:"script,omitempty"`

	Status *JobStatus `json:"status,omitempty"`

	JobModeSetting *IsapJobModeSettingDto `json:"job_mode_setting,omitempty"`

	JobOutputSetting *IsapJobOutputSetting `json:"job_output_setting,omitempty"`

	Environment *JobEnvironment `json:"environment,omitempty"`

	// UUID
	OutputTableId *string `json:"output_table_id,omitempty"`

	// 输出表ID列表
	OutputTableIds *[]string `json:"output_table_ids,omitempty"`

	// 输出表名称列表
	OutputTableNames *[]string `json:"output_table_names,omitempty"`

	// 发布状态: 只适用行管租户，不对外暴露参数
	PublishStatus *string `json:"publish_status,omitempty"`
}

func (o UpdateAlertRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlertRuleRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateAlertRuleRequestBody", string(data)}, " ")
}
