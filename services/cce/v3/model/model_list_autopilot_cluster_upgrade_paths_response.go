package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAutopilotClusterUpgradePathsResponse Response Object
type ListAutopilotClusterUpgradePathsResponse struct {

	// **参数解释：** API版本 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ApiVersion *string `json:"apiVersion,omitempty"`

	// **参数解释：** 资源类型 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Kind *string `json:"kind,omitempty"`

	Metadata *Metadata `json:"metadata,omitempty"`

	// **参数解释：** 升级路径集合 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	UpgradePaths   *[]UpgradePath `json:"upgradePaths,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListAutopilotClusterUpgradePathsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAutopilotClusterUpgradePathsResponse struct{}"
	}

	return strings.Join([]string{"ListAutopilotClusterUpgradePathsResponse", string(data)}, " ")
}
