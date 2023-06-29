package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeGraphReq This is a auto create Body Object
type UpgradeGraphReq struct {

	// 升级到的版本，必须大于当前图版本。
	UpgradeVersion string `json:"upgradeVersion"`

	// 是否强制升级。取值为true或false，默认为false。 - true：强制升级，会中断升级时已经在处理的任务，比如运行算法长任务，可能会造成少量请求失败。 - false：非强制升级，会等待已经运行的业务，升级过程可能较慢。
	ForceUpgrade *bool `json:"forceUpgrade,omitempty"`
}

func (o UpgradeGraphReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeGraphReq struct{}"
	}

	return strings.Join([]string{"UpgradeGraphReq", string(data)}, " ")
}
