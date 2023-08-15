package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DurationStrategies struct {

	// 升级策略 - minimized_interrupt_time，表示最短中断时间 - minimized_upgrade_time，最短升级时长
	Strategy string `json:"strategy"`

	// 升级时长，单位为分钟
	EstimatedUpgradeDuration int32 `json:"estimated_upgrade_duration"`
}

func (o DurationStrategies) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DurationStrategies struct{}"
	}

	return strings.Join([]string{"DurationStrategies", string(data)}, " ")
}
