package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeNodePoolSpecNodeTemplate 同步节点池模板参数
type UpgradeNodePoolSpecNodeTemplate struct {
	LifeCycle *NodeLifecycleConfig `json:"lifeCycle"`

	Login *Login `json:"login"`

	VolumeConfig *VolumeConfig `json:"volumeConfig,omitempty"`
}

func (o UpgradeNodePoolSpecNodeTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeNodePoolSpecNodeTemplate struct{}"
	}

	return strings.Join([]string{"UpgradeNodePoolSpecNodeTemplate", string(data)}, " ")
}
