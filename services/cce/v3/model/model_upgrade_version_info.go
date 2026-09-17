package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeVersionInfo **参数解释：** 版本信息 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type UpgradeVersionInfo struct {

	// **参数解释：** 正式版本号，如：v1.19.10 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Release *string `json:"release,omitempty"`

	// **参数解释：** 补丁版本号，如r0 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Patch *string `json:"patch,omitempty"`

	// **参数解释：** 推荐升级的目标补丁版本号，如r0 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	SuggestPatch *string `json:"suggestPatch,omitempty"`

	// **参数解释：** 升级目标版本集合 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	TargetVersions *[]string `json:"targetVersions,omitempty"`
}

func (o UpgradeVersionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeVersionInfo struct{}"
	}

	return strings.Join([]string{"UpgradeVersionInfo", string(data)}, " ")
}
