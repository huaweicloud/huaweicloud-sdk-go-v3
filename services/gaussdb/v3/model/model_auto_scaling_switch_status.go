package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AutoScalingSwitchStatus 开关状态。
type AutoScalingSwitchStatus struct {

	// 自动变配开关状态。  取值：  - ON：开启。 - OFF：关闭。
	ScalingSwitch *string `json:"scaling_switch,omitempty"`

	// 扩缩规格开关状态。  取值：  - ON：开启。 - OFF：关闭。
	FlavorSwitch *string `json:"flavor_switch,omitempty"`

	// 增删只读节点开关状态。  取值：  - ON：开启。 - OFF：关闭。
	ReadOnlySwitch *string `json:"read_only_switch,omitempty"`
}

func (o AutoScalingSwitchStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoScalingSwitchStatus struct{}"
	}

	return strings.Join([]string{"AutoScalingSwitchStatus", string(data)}, " ")
}
