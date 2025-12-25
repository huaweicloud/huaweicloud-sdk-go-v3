package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateAlertRuleRequestBody struct {

	// Alert rule name 告警规则名称
	AlertRuleName string `json:"alert_rule_name"`

	// Alert rule description 告警规则描述
	Description string `json:"description"`

	// directory 目录分组
	Directory *string `json:"directory,omitempty"`

	// Job Script 作业脚本
	Script *string `json:"script,omitempty"`

	Status *JobStatus `json:"status"`

	JobMode *IsapJobMode `json:"job_mode"`

	JobModeSetting *IsapJobModeSettingDto `json:"job_mode_setting,omitempty"`

	Environment *JobEnvironment `json:"environment"`

	JobOutputSetting *IsapJobOutputSetting `json:"job_output_setting,omitempty"`

	// UUID
	OutputTableId string `json:"output_table_id"`

	// 表名
	OutputTableName *string `json:"output_table_name,omitempty"`

	// 数量
	CuQuotaAmount float32 `json:"cu_quota_amount"`

	// 输出表id列表
	OutputTableIds []string `json:"output_table_ids"`

	// 输出表名称列表
	OutputTableNames []string `json:"output_table_names"`
}

func (o CreateAlertRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAlertRuleRequestBody struct{}"
	}

	return strings.Join([]string{"CreateAlertRuleRequestBody", string(data)}, " ")
}
