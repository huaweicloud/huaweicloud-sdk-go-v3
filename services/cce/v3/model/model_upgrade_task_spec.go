package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeTaskSpec **参数解释：** 升级任务属性 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type UpgradeTaskSpec struct {

	// **参数解释：** 升级前集群版本 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Version *string `json:"version,omitempty"`

	// **参数解释：** 升级的目标集群版本 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	TargetVersion *string `json:"targetVersion,omitempty"`

	// **参数解释：** 升级任务附属信息 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Items *interface{} `json:"items,omitempty"`
}

func (o UpgradeTaskSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeTaskSpec struct{}"
	}

	return strings.Join([]string{"UpgradeTaskSpec", string(data)}, " ")
}
