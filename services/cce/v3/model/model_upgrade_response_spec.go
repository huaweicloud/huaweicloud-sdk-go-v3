package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeResponseSpec **参数解释：** 升级任务元数据 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type UpgradeResponseSpec struct {
	ClusterUpgradeAction *ClusterUpgradeResponseAction `json:"clusterUpgradeAction,omitempty"`
}

func (o UpgradeResponseSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeResponseSpec struct{}"
	}

	return strings.Join([]string{"UpgradeResponseSpec", string(data)}, " ")
}
