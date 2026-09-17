package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeSpec **参数解释：** 集群升级配置详情。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type UpgradeSpec struct {
	ClusterUpgradeAction *ClusterUpgradeAction `json:"clusterUpgradeAction,omitempty"`
}

func (o UpgradeSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeSpec struct{}"
	}

	return strings.Join([]string{"UpgradeSpec", string(data)}, " ")
}
