package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 阶段运行信息
type StageRun struct {

	// 阶段ID
	Id *string `json:"id,omitempty"`

	// 阶段类型
	Category *string `json:"category,omitempty"`

	// 阶段名称
	Name *string `json:"name,omitempty"`

	// 唯一标识
	Identifier *string `json:"identifier,omitempty"`

	// 是否总是运行
	RunAlways *bool `json:"run_always,omitempty"`

	// 是否并行
	Parallel *string `json:"parallel,omitempty"`

	// 是否选中
	IsSelect *bool `json:"is_select,omitempty"`

	// 序列号
	Sequence *int32 `json:"sequence,omitempty"`

	// 依赖
	DependsOn *[]string `json:"depends_on,omitempty"`

	// 运行条件
	Condition *string `json:"condition,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 开始时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 结束时间
	EndTime *int64 `json:"end_time,omitempty"`

	// 阶段准入
	Pre *[]StepRun `json:"pre,omitempty"`

	// 阶段准出
	Post *[]StepRun `json:"post,omitempty"`

	// 任务
	Jobs *[]JobRun `json:"jobs,omitempty"`
}

func (o StageRun) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StageRun struct{}"
	}

	return strings.Join([]string{"StageRun", string(data)}, " ")
}
