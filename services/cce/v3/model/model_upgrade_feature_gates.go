package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeFeatureGates **参数解释：** 集群升级特性开关 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type UpgradeFeatureGates struct {

	// **参数解释：** 集群升级Console界面是否支持V4版本，该字段一般由CCE Console使用。 **约束限制：** 不涉及 **取值范围：** - true：支持V4版本 - false：不支持V4版本  **默认取值：** 不涉及
	SupportUpgradePageV4 *bool `json:"supportUpgradePageV4,omitempty"`
}

func (o UpgradeFeatureGates) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeFeatureGates struct{}"
	}

	return strings.Join([]string{"UpgradeFeatureGates", string(data)}, " ")
}
