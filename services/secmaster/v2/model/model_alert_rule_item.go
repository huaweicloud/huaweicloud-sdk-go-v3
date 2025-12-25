package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlertRuleItem struct {

	// UUID
	AlertRuleId string `json:"alert_rule_id"`

	// Alert rule name 告警规则名称
	AlertRuleName string `json:"alert_rule_name"`

	// Job Script 作业脚本
	Script *string `json:"script,omitempty"`

	Status *JobStatus `json:"status"`

	// directory 目录分组
	Directory *string `json:"directory,omitempty"`

	// Alert rule description 告警规则描述
	Description *string `json:"description,omitempty"`

	JobMode *IsapJobMode `json:"job_mode"`

	ProcessStatus *JobProcessStatus `json:"process_status"`

	ProcessError *AlertRuleProcessError `json:"process_error,omitempty"`

	Environment *JobEnvironment `json:"environment"`

	// UUID
	OutputTableId string `json:"output_table_id"`

	// 表名称
	OutputTableName string `json:"output_table_name"`

	// 输出表ID列表
	OutputTableIds []string `json:"output_table_ids"`

	// 输出表名称列表
	OutputTableNames []string `json:"output_table_names"`

	// 创建者
	CreateBy string `json:"create_by"`

	// 毫秒时间戳
	CreateTime int64 `json:"create_time"`

	// 更新者
	UpdateBy string `json:"update_by"`

	// 毫秒时间戳
	UpdateTime int64 `json:"update_time"`

	// 毫秒时间戳
	DeleteTime *int64 `json:"delete_time,omitempty"`
}

func (o AlertRuleItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlertRuleItem struct{}"
	}

	return strings.Join([]string{"AlertRuleItem", string(data)}, " ")
}
