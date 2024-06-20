package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PipelineConcurrencyMgmt 流水线并发管理
type PipelineConcurrencyMgmt struct {

	// 流水线id
	PipelineId *string `json:"pipeline_id,omitempty"`

	// 并行数量
	ConcurrencyNumber *int32 `json:"concurrency_number,omitempty"`

	// 超出情况下策略
	ExceedAction *string `json:"exceed_action,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 是否有效
	Enable *bool `json:"enable,omitempty"`
}

func (o PipelineConcurrencyMgmt) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineConcurrencyMgmt struct{}"
	}

	return strings.Join([]string{"PipelineConcurrencyMgmt", string(data)}, " ")
}
