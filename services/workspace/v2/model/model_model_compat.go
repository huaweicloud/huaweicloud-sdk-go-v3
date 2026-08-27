package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModelCompat 模型兼容性配置。
type ModelCompat struct {

	// 是否支持使用量流式传输。
	SupportsUsageStreaming *bool `json:"supports_usage_streaming,omitempty"`

	// 是否支持开发者角色。
	SupportsDeveloperRole *bool `json:"supports_developer_role,omitempty"`
}

func (o ModelCompat) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModelCompat struct{}"
	}

	return strings.Join([]string{"ModelCompat", string(data)}, " ")
}
