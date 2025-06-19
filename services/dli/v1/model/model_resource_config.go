package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceConfig 资源配置。
type ResourceConfig struct {

	// 该参数用于设置单个TaskManager可以提供的并行任务数量。每个Task Slot可以并行执行一个任务。增加 Task Slots 可以提高 TaskManager 的并行处理能力，但也会增加资源消耗。Task Slots的数量与TaskManager的CPU数相关联，因为每个CPU可以提供一个Task Slot。单TM Slot默认值为1。最小并行数不能小于1。
	MaxSlot *int32 `json:"max_slot,omitempty"`

	// 作业的并行数，指作业中各个算子的并行执行的子任务的数量，算子的子任务数就是其对应算子的并行度。默认值为“1”。
	ParallelNumber *int32 `json:"parallel_number,omitempty"`

	JobmanagerResourceSpec *ResourceSpec `json:"jobmanager_resource_spec,omitempty"`

	TaskmanagerResourceSpec *ResourceSpec `json:"taskmanager_resource_spec,omitempty"`
}

func (o ResourceConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceConfig struct{}"
	}

	return strings.Join([]string{"ResourceConfig", string(data)}, " ")
}
