package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EnvironmentMetadata struct {

	// 附加信息。
	Annotations *interface{} `json:"annotations,omitempty"`

	// 环境id。
	Id *string `json:"id,omitempty"`

	// 环境名称。
	Name *string `json:"name,omitempty"`

	// 环境状态。
	Status *string `json:"status,omitempty"`

	// 环境类型。
	Type *string `json:"type,omitempty"`

	// 任务id。
	JobId *string `json:"job_id,omitempty"`

	// 创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间。
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (o EnvironmentMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnvironmentMetadata struct{}"
	}

	return strings.Join([]string{"EnvironmentMetadata", string(data)}, " ")
}
