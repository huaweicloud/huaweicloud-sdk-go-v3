package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskEntity 任务实体
type TaskEntity struct {

	// 任务ID
	TaskId string `json:"task_id"`

	TaskType *TaskTypeEnum `json:"task_type"`

	TaskStatus *TaskStatusEnum `json:"task_status"`

	// 资源名称
	ResourceName string `json:"resource_name"`

	// 资源ID
	ResourceId string `json:"resource_id"`

	ResourceType *ResourceTypeEnum `json:"resource_type"`

	// 任务归属的区域ID
	RegionId string `json:"region_id"`

	// 存储库名称
	VaultName *string `json:"vault_name,omitempty"`

	// 存储库ID
	VaultId *string `json:"vault_id,omitempty"`

	// 任务开始时间
	TaskStartTime string `json:"task_start_time"`

	// 任务结束时间
	TaskEndTime string `json:"task_end_time"`
}

func (o TaskEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskEntity struct{}"
	}

	return strings.Join([]string{"TaskEntity", string(data)}, " ")
}
