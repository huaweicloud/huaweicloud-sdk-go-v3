package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SignatureExecutionTask struct {

	// 执行任务ID
	Id *int64 `json:"id,omitempty"`

	// 执行ID
	ExecutionId *int64 `json:"execution_id,omitempty"`

	// jobId
	JobId *string `json:"job_id,omitempty"`

	// 执行任务状态
	Status *string `json:"status,omitempty"`

	// 执行任务状态信息
	StatusText *string `json:"status_text,omitempty"`

	// 命名空间
	Namespace *string `json:"namespace,omitempty"`

	// 镜像仓库名称
	Repository *string `json:"repository,omitempty"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
}

func (o SignatureExecutionTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SignatureExecutionTask struct{}"
	}

	return strings.Join([]string{"SignatureExecutionTask", string(data)}, " ")
}
