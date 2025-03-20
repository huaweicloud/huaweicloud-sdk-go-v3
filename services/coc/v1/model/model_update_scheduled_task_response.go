package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateScheduledTaskResponse Response Object
type UpdateScheduledTaskResponse struct {

	// 任务ID
	Id *string `json:"id,omitempty"`

	// 企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 任务名称
	Name *string `json:"name,omitempty"`

	// 委托名称
	AgencyName *string `json:"agency_name,omitempty"`

	TriggerTime *TriggerTime `json:"trigger_time,omitempty"`

	// 版本号
	VersionNo *string `json:"version_no,omitempty"`

	// 任务类型
	TaskType *interface{} `json:"task_type,omitempty"`

	// 关联任务ID（脚本ID/作业ID）
	AssociatedTaskId *string `json:"associated_task_id,omitempty"`

	// 关联任务名称（脚本名称/作业名称）
	AssociatedTaskName *string `json:"associated_task_name,omitempty"`

	// 关联任务名称(英文)（脚本名称/作业名称）
	AssociatedTaskNameEn *string `json:"associated_task_name_en,omitempty"`

	// 关联任务的类型
	AssociatedTaskType *string `json:"associated_task_type,omitempty"`

	// 作业实例模式
	RunbookInstanceMode *string `json:"runbook_instance_mode,omitempty"`

	// 风险等级
	RiskLevel *string `json:"risk_level,omitempty"`

	// 执行参数，值为json串
	InputParam *string `json:"input_param,omitempty"`

	// 是否开启入库人工审核
	EnableApprove *bool `json:"enable_approve,omitempty"`

	ReviewerNotification *MessageNotification `json:"reviewer_notification,omitempty"`

	// 创建人昵称
	CreatedUserName *string `json:"created_user_name,omitempty"`

	// 审核人昵称
	ReviewerUserName *string `json:"reviewer_user_name,omitempty"`

	// 审批状态
	ApproveStatus *interface{} `json:"approve_status,omitempty"`

	// 审批意见
	ApproveComments *string `json:"approve_comments,omitempty"`

	// 目标节点json串
	TargetInstances *string `json:"target_instances,omitempty"`

	// 是否启用消息通知
	EnableMessageNotification *bool `json:"enable_message_notification,omitempty"`

	MessageNotification *MessageNotification `json:"message_notification,omitempty"`
	HttpStatusCode      int                  `json:"-"`
}

func (o UpdateScheduledTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScheduledTaskResponse struct{}"
	}

	return strings.Join([]string{"UpdateScheduledTaskResponse", string(data)}, " ")
}
