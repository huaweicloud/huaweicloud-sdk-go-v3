package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScheduledTaskBasicInfo 定时运维基本信息
type ScheduledTaskBasicInfo struct {

	// 任务ID
	Id *string `json:"id,omitempty"`

	// 企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 任务名称
	Name *string `json:"name,omitempty"`

	// 定时类型
	ScheduledType *string `json:"scheduled_type,omitempty"`

	// 引用任务类型
	TaskType *string `json:"task_type,omitempty"`

	// 关联的任务类型
	AssociatedTaskType *string `json:"associated_task_type,omitempty"`

	// 风险等级
	RiskLevel *string `json:"risk_level,omitempty"`

	// 创建人
	CreatedBy *string `json:"created_by,omitempty"`

	// 修改人
	UpdateBy *string `json:"update_by,omitempty"`

	// 创建人昵称
	CreatedUserName *string `json:"created_user_name,omitempty"`

	// 审批人
	Reviewer *string `json:"reviewer,omitempty"`

	// 审批人昵称
	ReviewerUserName *string `json:"reviewer_user_name,omitempty"`

	// 审批状态
	ApproveStatus *interface{} `json:"approve_status,omitempty"`

	// 最近执行时间时间戳
	LastExecutionTime *int64 `json:"last_execution_time,omitempty"`

	// 最近执行状态
	LastExecutionStatus *string `json:"last_execution_status,omitempty"`

	// 执行次数
	ExecutionCount *int32 `json:"execution_count,omitempty"`

	// 是否启用
	Enabled *bool `json:"enabled,omitempty"`

	// 创建时间
	CreatedTime *int64 `json:"created_time,omitempty"`

	// 更新时间
	ModifiedTime *int64 `json:"modified_time,omitempty"`

	// 区域
	RegionId *string `json:"region_id,omitempty"`

	// 脚本/作业名称
	AssociatedTaskName *string `json:"associated_task_name,omitempty"`

	// 脚本/作业名称(英文)
	AssociatedTaskNameEn *string `json:"associated_task_name_en,omitempty"`
}

func (o ScheduledTaskBasicInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScheduledTaskBasicInfo struct{}"
	}

	return strings.Join([]string{"ScheduledTaskBasicInfo", string(data)}, " ")
}
