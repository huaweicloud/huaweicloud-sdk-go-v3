package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlinkResourceConfig Flink 作业的资源配置。
type FlinkResourceConfig struct {

	// 最大的 slot 数。
	MaxSlot *int32 `json:"max_slot,omitempty"`

	// 用户设置的作业并行数目。默认值为1。
	ParallelNumber *int32 `json:"parallel_number,omitempty"`

	JobmanagerResourceSpec *ResourceSpec `json:"jobmanager_resource_spec,omitempty"`

	TaskmanagerResourceSpec *ResourceSpec `json:"taskmanager_resource_spec,omitempty"`
}

func (o FlinkResourceConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkResourceConfig struct{}"
	}

	return strings.Join([]string{"FlinkResourceConfig", string(data)}, " ")
}
