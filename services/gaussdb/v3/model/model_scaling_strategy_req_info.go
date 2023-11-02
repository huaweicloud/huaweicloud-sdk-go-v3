package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScalingStrategyReqInfo 自动变配策略。
type ScalingStrategyReqInfo struct {

	// 扩缩规格开关。 - ON：开启。 - OFF：关闭。
	FlavorSwitch string `json:"flavor_switch"`

	// 增删只读节点开关。自动增删只读节点功能需要先[开启数据库代理](https://support.huaweicloud.com/api-gaussdbformysql/CreateGaussMySqlProxy.html)，且只能有一个代理。 - ON：开启。 - OFF：关闭。
	ReadOnlySwitch string `json:"read_only_switch"`
}

func (o ScalingStrategyReqInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScalingStrategyReqInfo struct{}"
	}

	return strings.Join([]string{"ScalingStrategyReqInfo", string(data)}, " ")
}
