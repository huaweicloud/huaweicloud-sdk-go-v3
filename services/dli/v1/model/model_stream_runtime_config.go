package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StreamRuntimeConfig 流作业运行时配置。
type StreamRuntimeConfig struct {

	// 临时文件存储 URI，作业运行时产生的临时文件存储的 OBS 路径。（当前不支持配置）
	StagingUri *string `json:"staging_uri,omitempty"`

	Logging *StreamLoggingConfig `json:"logging,omitempty"`

	Properties *Properties `json:"properties,omitempty"`

	FlinkRuntimeConfig *FlinkRuntimeConfig `json:"flink_runtime_config,omitempty"`
}

func (o StreamRuntimeConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StreamRuntimeConfig struct{}"
	}

	return strings.Join([]string{"StreamRuntimeConfig", string(data)}, " ")
}
