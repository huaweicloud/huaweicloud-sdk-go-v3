package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NextflowTaskListDto struct {

	// 子任务id
	TaskId *string `json:"task_id,omitempty"`

	// 流程名称
	Process *string `json:"process,omitempty"`

	// 子任务标识符
	Tag *string `json:"tag,omitempty"`

	// 哈希值
	Hash *string `json:"hash,omitempty"`

	// 子任务状态
	Status *string `json:"status,omitempty"`

	// 容器名称
	Container *string `json:"container,omitempty"`

	// pod名称
	PodName *string `json:"pod_name,omitempty"`

	// 提交时间
	Submit *string `json:"submit,omitempty"`

	// 完成时间
	Complete *string `json:"complete,omitempty"`

	// 总时间
	Duration *int64 `json:"duration,omitempty"`

	// 实际运行时间
	Realtime *int64 `json:"realtime,omitempty"`

	// cpu使用率
	CpuPercent *float64 `json:"cpu_percent,omitempty"`

	// 内存使用率
	MemPercent *float64 `json:"mem_percent,omitempty"`
}

func (o NextflowTaskListDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NextflowTaskListDto struct{}"
	}

	return strings.Join([]string{"NextflowTaskListDto", string(data)}, " ")
}
