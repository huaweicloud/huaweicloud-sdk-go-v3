package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeInfoSpec **参数解释：** 升级配置相关信息 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type UpgradeInfoSpec struct {
	LastUpgradeInfo *UpgradeInfoStatus `json:"lastUpgradeInfo,omitempty"`

	VersionInfo *UpgradeVersionInfo `json:"versionInfo,omitempty"`

	UpgradeFeatureGates *UpgradeFeatureGates `json:"upgradeFeatureGates,omitempty"`
}

func (o UpgradeInfoSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeInfoSpec struct{}"
	}

	return strings.Join([]string{"UpgradeInfoSpec", string(data)}, " ")
}
