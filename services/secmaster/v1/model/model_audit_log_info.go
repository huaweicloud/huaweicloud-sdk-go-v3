package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AuditLogInfo 剧本实例审计日志信息
type AuditLogInfo struct {

	// 实例类型（AOP_WORKFLOW--流程, SCRIPT--脚本, PLAYBOOK--剧本）
	InstanceType *string `json:"instance_type,omitempty"`

	// 流程ID
	ActionId *string `json:"action_id,omitempty"`

	// 流程名称
	ActionName *string `json:"action_name,omitempty"`

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 父节点实例ID
	ParentInstanceId *string `json:"parent_instance_id,omitempty"`

	// 日志级别
	LogLevel *string `json:"log_level,omitempty"`

	// 输入
	Input *string `json:"input,omitempty"`

	// 输出
	Output *string `json:"output,omitempty"`

	// 错误信息
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 状态。(RUNNING--运行中、FINISHED--成功、FAILED--失败、RETRYING--重试中、TERMINATING--终止中、TERMINATED--已终止)
	Status *string `json:"status,omitempty"`

	// 触发类型. TIMER--定时触发, EVENT--事件触发
	TriggerType *string `json:"trigger_type,omitempty"`
}

func (o AuditLogInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuditLogInfo struct{}"
	}

	return strings.Join([]string{"AuditLogInfo", string(data)}, " ")
}
