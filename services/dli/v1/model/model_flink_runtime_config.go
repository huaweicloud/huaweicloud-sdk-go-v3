package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlinkRuntimeConfig Flink 作业运行时配置。
type FlinkRuntimeConfig struct {

	// Flink 版本。仅支持 Flink 1.15及以上版本。
	Version *string `json:"version,omitempty"`

	RestoreStrategy *FlinkRestoreStrategy `json:"restore_strategy,omitempty"`

	ResourceConfig *FlinkResourceConfig `json:"resource_config,omitempty"`
}

func (o FlinkRuntimeConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkRuntimeConfig struct{}"
	}

	return strings.Join([]string{"FlinkRuntimeConfig", string(data)}, " ")
}
