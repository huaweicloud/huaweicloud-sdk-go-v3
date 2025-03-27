package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SpeedLimitlJson 速率参数
type SpeedLimitlJson struct {

	// 时间段开始时间，格式：XX:XX。
	Start string `json:"start"`

	// 时间段结束时间，格式：XX:XX。
	End string `json:"end"`

	// 时间段的速率，0-1000的整数，单位：Mbit/s。
	Speed int32 `json:"speed"`

	// 停止迁移的超速阈值。 是一个迁移速率的保护机制，超出该阈值会停止任务。它主要用于控制迁移过程中资源（特别是网络带宽）的消耗，确保系统的整体性能不受单一迁移任务影响 单位是百分比
	OverSpeedThreshold *float64 `json:"over_speed_threshold,omitempty"`
}

func (o SpeedLimitlJson) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SpeedLimitlJson struct{}"
	}

	return strings.Join([]string{"SpeedLimitlJson", string(data)}, " ")
}
