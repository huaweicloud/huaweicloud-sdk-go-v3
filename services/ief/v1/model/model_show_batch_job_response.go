package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBatchJobResponse Response Object
type ShowBatchJobResponse struct {

	// 批量处理作业ID
	JobId *string `json:"job_id,omitempty"`

	// 批量处理作业名称
	JobName *string `json:"job_name,omitempty"`

	// 批量作业类型，支持以下选项： - node_upgrade： 节点升级 - deployment_deploy：应用部署 - deployment_upgrade：应用升级
	JobType *string `json:"job_type,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 创建时间戳
	CreatedAt *int32 `json:"created_at,omitempty"`

	// 执行状态
	Status *string `json:"status,omitempty"`

	// 任务总数
	TaskTotalCount *int32 `json:"task_total_count,omitempty"`

	// 任务项执行成功数
	TaskFailedCount *int32 `json:"task_failed_count,omitempty"`

	// 任务项执行失败数
	TaskSuccessCount *int32 `json:"task_success_count,omitempty"`

	// 批量作业对象类型，支持如下选项： - node：边缘节点 - node_group：边缘节点组 - deployment：边缘应用
	TargetType *string `json:"target_type,omitempty"`

	// 批量作业内容，仅在批量应用部署和批量应用升级时需要填写，填入的内容为：使用json结构体编写的创建应用部署接口请求体deployment参数，并将其转换为字符串
	TaskData *string `json:"task_data,omitempty"`

	// 任务项详情
	Tasks *[]Task `json:"tasks,omitempty"`

	// 批量处理对象详情
	Targets *[]Target `json:"targets,omitempty"`

	// 状态更新时间戳
	StatusLastUpdatedAt *int32 `json:"status_last_updated_at,omitempty"`
	HttpStatusCode      int    `json:"-"`
}

func (o ShowBatchJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBatchJobResponse struct{}"
	}

	return strings.Join([]string{"ShowBatchJobResponse", string(data)}, " ")
}
