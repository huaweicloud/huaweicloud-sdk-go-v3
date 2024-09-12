package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AudioDriveActionConfig 语音驱动时的动作标签配置
type AudioDriveActionConfig struct {

	// 动作标签
	ActionTag string `json:"action_tag"`

	// 动作名称
	ActionName *string `json:"action_name,omitempty"`

	// 动作开始时间
	ActionStartTime float32 `json:"action_start_time"`
}

func (o AudioDriveActionConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AudioDriveActionConfig struct{}"
	}

	return strings.Join([]string{"AudioDriveActionConfig", string(data)}, " ")
}
