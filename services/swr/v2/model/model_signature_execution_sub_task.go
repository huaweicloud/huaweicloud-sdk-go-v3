package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SignatureExecutionSubTask struct {

	// 子任务ID
	Id *int64 `json:"id,omitempty"`

	// 子任务详情
	JobId *string `json:"job_id,omitempty"`

	// 命名空间
	Namespace *string `json:"namespace,omitempty"`

	// 镜像仓库
	Repository *string `json:"repository,omitempty"`

	// 镜像版本
	Tags *string `json:"tags,omitempty"`

	// 镜像层sha256
	Digest *string `json:"digest,omitempty"`

	// 子任务状态
	Status *string `json:"status,omitempty"`

	// 状态详情信息
	StatusText *string `json:"status_text,omitempty"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
}

func (o SignatureExecutionSubTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SignatureExecutionSubTask struct{}"
	}

	return strings.Join([]string{"SignatureExecutionSubTask", string(data)}, " ")
}
