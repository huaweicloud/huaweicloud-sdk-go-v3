package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MetadataTask 元数据迁移任务。
type MetadataTask struct {

	// 元数据迁移任务ID。
	Id *string `json:"id,omitempty"`

	// 元数据迁移任务名称。
	Name *string `json:"name,omitempty"`

	// 元数据迁移任务开始时间。
	StartDate *string `json:"start_date,omitempty"`

	// 元数据迁移任务状态。
	Status *string `json:"status,omitempty"`

	// 元数据迁移类型。
	Type *string `json:"type,omitempty"`
}

func (o MetadataTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetadataTask struct{}"
	}

	return strings.Join([]string{"MetadataTask", string(data)}, " ")
}
