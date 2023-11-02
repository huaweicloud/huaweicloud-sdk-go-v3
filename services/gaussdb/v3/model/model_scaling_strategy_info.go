package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScalingStrategyInfo 自动变配策略。
type ScalingStrategyInfo struct {

	// 扩缩规格开关。 - ON：开启。 - OFF：关闭。
	FlavorSwitch *string `json:"flavor_switch,omitempty"`

	// 增删只读节点开关。 - ON：开启。 - OFF：关闭。
	ReadOnlySwitch *string `json:"read_only_switch,omitempty"`
}

func (o ScalingStrategyInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScalingStrategyInfo struct{}"
	}

	return strings.Join([]string{"ScalingStrategyInfo", string(data)}, " ")
}
