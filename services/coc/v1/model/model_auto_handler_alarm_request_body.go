package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AutoHandlerAlarmRequestBody 创建任务请求体
type AutoHandlerAlarmRequestBody struct {

	// 任务类型
	TaskType string `json:"task_type"`

	// 关联任务ID（脚本ID/作业ID）
	AssociatedTaskId string `json:"associated_task_id"`

	// 关联任务的类型
	AssociatedTaskType string `json:"associated_task_type"`

	// 关联任务名称（脚本名称/作业名称）
	AssociatedTaskName string `json:"associated_task_name"`

	// 企业项目ID
	AssociatedTaskEnterpriseProjectId *string `json:"associated_task_enterprise_project_id,omitempty"`

	// 作业实例模式
	RunbookInstanceMode *string `json:"runbook_instance_mode,omitempty"`

	// 执行参数，值为json串
	InputParam map[string]string `json:"input_param"`

	// 目标实例信息
	TargetInstances *[]AlarmScheduleInstance `json:"target_instances,omitempty"`

	// 区域ID
	RegionId *string `json:"region_id,omitempty"`

	SubTaskInfo *SubTaskInfoDto `json:"sub_task_info,omitempty"`
}

func (o AutoHandlerAlarmRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoHandlerAlarmRequestBody struct{}"
	}

	return strings.Join([]string{"AutoHandlerAlarmRequestBody", string(data)}, " ")
}
