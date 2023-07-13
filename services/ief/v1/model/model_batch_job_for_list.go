package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchJobForList 批量处理作业详情
type BatchJobForList struct {

	// 批量处理作业ID
	JobId *string `json:"job_id,omitempty"`

	// 批量处理作业名称
	JobName *string `json:"job_name,omitempty"`

	// 批量处理作业类型，支持以下选项： - node_upgrade： 节点升级 - deployment_deploy：应用部署 - deployment_upgrade：应用升级
	JobType *string `json:"job_type,omitempty"`

	// 创建时间戳
	CreatedAt *int32 `json:"created_at,omitempty"`

	// 执行状态
	Status *string `json:"status,omitempty"`

	// 任务总数
	TaskTotalCount *int32 `json:"task_total_count,omitempty"`

	// 任务项执行成功数
	TaskSuccessCount *int32 `json:"task_success_count,omitempty"`

	// 任务项执行失败数
	TaskFailedCount *int32 `json:"task_failed_count,omitempty"`

	// 状态更新时间戳
	StatusLastUpdatedAt *int32 `json:"status_last_updated_at,omitempty"`

	// 任务描述
	Description *string `json:"description,omitempty"`
}

func (o BatchJobForList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchJobForList struct{}"
	}

	return strings.Join([]string{"BatchJobForList", string(data)}, " ")
}
